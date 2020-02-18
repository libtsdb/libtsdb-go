package gorilla_test

import (
	"encoding/binary"
	"io"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// a quick hack implementation for the algorithm described in gorilla paper

// bits is based on https://github.com/dgryski/go-tsz/blob/master/bstream.go
// it allows you to read/write individual bit(s)
type bits struct {
	buf []byte // underlying bytes
	i   int    // index of last byte
	// NOTE: the reason we use remain instead of used because it makes appending bits easier
	// byte = byte | 1 << (remain - 1)
	remain uint8
}

func newBits() *bits {
	// 0 byte and 0 bits, i is -1 so the grow logic works ...
	return &bits{buf: make([]byte, 0), i: -1, remain: 0}
}

func (b *bits) writeBit(bit bool) {
	if b.remain == 0 {
		b.buf = append(b.buf, 0)
		b.remain = 8
		b.i++
	}
	if bit {
		b.buf[b.i] |= 1 << (b.remain - 1)
	}
	b.remain--
}

func (b *bits) writeByte(byt byte) {
	// fast path, previous write are aligned to byte boundary
	if b.remain == 0 {
		b.buf = append(b.buf, byt)
		b.i++
		return
	}

	// e.g. b.remain = 6
	//       [0, 1, 2, 3, 4, 5,  6, 7]
	// [0, 1, 2, 3, 4, 5, 6, 7] [0, 1, 2, 3, 4, 5, 6, 7]
	b.buf[b.i] |= byt >> (8 - b.remain)
	b.buf = append(b.buf, 0)
	b.i++
	b.buf[b.i] |= byt << b.remain
	// no need to update b.remain, it's the same
}

func (b *bits) writeBits(u uint64, n uint) {
	u <<= 64 - n
	for n >= 8 {
		byt := byte(u >> 56)
		b.writeByte(byt)
		u <<= 8
		n -= 8
	}

	for n > 0 {
		b.writeBit((u >> 63) == 1)
		u <<= 1
		n--
	}
}

// bitsReader follows the modified version in prometheus where the underlying byte is not modified when read
type bitsReader struct {
	buf      []byte
	original []byte
	remain   uint8
}

func newReader(b []byte) *bitsReader {
	remain := 0
	if len(b) != 0 {
		remain = 8
	}
	return &bitsReader{
		buf:      b,
		original: b,
		remain:   uint8(remain),
	}
}

func (r *bitsReader) reset() {
	r.buf = r.original
}

func (r *bitsReader) readBit() (bool, error) {
	if r.remain == 0 {
		if len(r.buf) < 1 {
			return false, io.EOF
		}
		r.buf = r.buf[1:]
		r.remain = 8
	}

	b := r.buf[0] & 0b1000_0000
	r.remain--
	return b != 0, nil
}

func (r *bitsReader) readByte() (byte, error) {
	if r.remain == 0 {
		if len(r.buf) < 1 {
			return 0, io.EOF
		}
		r.buf = r.buf[1:]
		return r.buf[0], nil
	}

	// We need to read a byte that cross two bytes
	if len(r.buf) < 1 {
		return 0, io.EOF
	}

	byt := r.buf[0] << (8 - r.remain)
	r.buf = r.buf[1:]
	byt |= r.buf[0] >> r.remain
	return byt, nil
}

func (r *bitsReader) readBits(n int) (uint64, error) {
	// TODO: return error when n is larger than 64?
	var u uint64
	for n >= 8 {
		byt, err := r.readByte()
		if err != nil {
			return 0, err
		}
		u = (u << 8) | uint64(byt)
		n -= 8
	}

	if n == 0 {
		return u, nil
	}

	if n > int(r.remain) {
		u = (u << r.remain) | uint64((r.buf[0]<<(8-r.remain))>>(8-r.remain))
		n -= int(r.remain)
		r.buf = r.buf[1:]
		if len(r.buf) == 0 {
			return 0, io.EOF
		}
		// FIXED: took me a while to find it ...
		r.remain = 8
	}

	u = (u << n) | uint64((r.buf[0]<<(8-r.remain))>>(8-n))
	r.remain -= uint8(n)
	return u, nil
}

func TestBitsWriter(t *testing.T) {
	t.Run("writeBit", func(t *testing.T) {
		bs := newBits()
		for i := 0; i < 8; i++ {
			bs.writeBit(true)
		}
		assert.Equal(t, bs.remain, uint8(0))
		assert.Equal(t, bs.buf[0], byte(0b1111_1111))
		bs.writeByte(8)
		assert.Equal(t, bs.buf[1], byte(8))
		bs.writeBit(true)
		bs.writeByte(1)
		assert.Equal(t, bs.buf[2], byte(0b1000_0000))
		assert.Equal(t, bs.buf[3], byte(0b1000_0000))
	})

	t.Run("writeBits", func(t *testing.T) {
		bs := newBits()
		bs.writeBits(20, 32)
		assert.Equal(t, bs.buf[0], byte(0))
		assert.Equal(t, bs.buf[1], byte(0))
		assert.Equal(t, bs.buf[2], byte(0))
		assert.Equal(t, bs.buf[3], byte(20))
		assert.Equal(t, len(bs.buf), 4)
		assert.Equal(t, bs.remain, uint8(0))
		assert.Equal(t, bs.i, 3)
	})

}

func TestBitsReader(t *testing.T) {
	w := newBits()
	w.writeBit(true)
	w.writeByte(1)
	cp := make([]byte, len(w.buf))
	copy(cp, w.buf)
	r := newReader(w.buf)
	b1, e1 := r.readBit()
	assert.Nil(t, e1)
	assert.Equal(t, true, b1)
	byt2, e2 := r.readByte()
	assert.Nil(t, e2)
	assert.Equal(t, uint8(1), byt2)
	assert.Equal(t, cp, w.buf)
}

// encoder encodes time stream, i.e. it does not mix value into same stream
type encoder struct {
	bs       bits
	start    uint64
	prevTime uint64
	delta    uint64
}

func newEncoder(start uint64) *encoder {
	bs := newBits()
	bs.writeBits(start, 64)
	return &encoder{
		bs:       *bs,
		start:    start,
		prevTime: 0,
	}
}

func (e *encoder) write(tm uint64) {
	// first value since start, write using delta
	if e.prevTime == 0 {
		delta := tm - e.start
		e.prevTime = tm
		e.bs.writeBits(delta, 14)
		e.delta = delta
		return
	}

	// TODO: delta is positive if time comes in order, dod can be negative because interval
	// double delta
	delta := tm - e.prevTime
	dod := int64(delta - e.delta)
	log.Printf("dod %d", dod)
	e.delta = delta
	switch {
	case dod == 0:
		e.bs.writeBit(false)
	case dod <= 64 && dod >= -63:
		log.Printf("dod write before %d %v", dod, e.bs.buf)
		e.bs.writeBits(0b10, 2)
		log.Printf("dod write middle %d %v", dod, e.bs.buf)
		e.bs.writeBits(uint64(dod), 7)
		log.Printf("dod write after %d %v", dod, e.bs.buf)
	case dod <= 256 && dod > -255:
		e.bs.writeBits(0b110, 3)
		e.bs.writeBits(uint64(dod), 9)
	case dod <= 2048 && dod > -2047:
		e.bs.writeBits(0b1110, 4)
		e.bs.writeBits(uint64(dod), 12)
	default:
		e.bs.writeBits(0b1111, 4)
		e.bs.writeBits(uint64(dod), 32)
	}
	e.prevTime = tm
}

func TestDoubleDelta(t *testing.T) {
	// Figure 2 in paper, start is aligned to 2 hour window
	start := mtime("2015-03-24T02:00:00Z")
	t1 := mtime("2015-03-24T02:01:02Z")
	t2 := mtime("2015-03-24T02:02:02Z")
	//t3 := mtime("2015-03-24T02:03:02Z")
	enc := newEncoder(start)
	enc.write(t1)
	enc.write(t2)
	//enc.write(t3)
	r := newReader(enc.bs.buf)
	log.Printf("%v", r.buf)
	// first 64 bytes is the header
	var b8 [8]byte
	binary.BigEndian.PutUint64(b8[:], start)
	for i := 0; i < 8; i++ {
		b, err := r.readByte()
		assert.Nil(t, err)
		assert.Equal(t, b8[i], b)
	}
	log.Printf("b8 %v buf %v", b8, r.buf)
	// the next 14 bits is the first time using delta
	// 62 is 111110, first 8 bits is empty, next 6 bits is the value
	delta, _ := r.readBits(14)
	assert.Equal(t, uint64(0b1111_10), delta)
	assert.Equal(t, t1-start, delta)
	// the first double delta encoded value, dict is 10, value is -2
	dict, _ := r.readBits(2)
	assert.Equal(t, uint64(0b10), dict)
	// FIXME: got 0 while it should be -2, is it because it's the last byte?
	log.Printf("%v", r.buf)
	v, err := r.readBits(7)
	log.Printf("%v", r.buf)
	assert.Nil(t, err)
	log.Print(v)
	assert.Equal(t, int64(-2), uint2int(v, 7))
}

// TODO: why this work ...
// https://stackoverflow.com/questions/4975340/int-to-unsigned-int-conversion ?
func uint2int(v uint64, sz int) int64 {
	if v > (1 << (sz - 1)) {
		// or something
		v = v - (1 << sz)
	}
	return int64(v)
}

func subu64(a, b uint64) int64 {
	return int64(a - b)
}

func TestUint64(t *testing.T) {
	// ./gorilla_test.go:164:23: constant -1 overflows uint64
	//a := int64(uint64(1) - uint64(2))
	//t.Log(a)
	// TODO: does this unsigned subtraction produce signed integer work in other languages?
	assert.Equal(t, subu64(1, 2), int64(-1))

	// cast is using the same bytes, but
	a := int64(-1)
	b := uint64(a)
	c := int64(b)
	t.Log(a, b, c) // -1 18446744073709551615 -1

	bs := newBits()
	three := uint64(3)
	five := uint64(5)
	a = int64(three - five)
	bs.writeBits(uint64(a), 7)
	t.Logf("%v", bs.buf)
	br := newReader(bs.buf)
	v, _ := br.readBits(7)
	v2 := uint2int(v, 7)
	t.Logf("%d %d %d", v, int64(v), v2)
	assert.Equal(t, int64(-2), v2)
}

func TestReadNegative(t *testing.T)  {
	start := mtime("2015-03-24T02:00:00Z")
	t1 := mtime("2015-03-24T02:01:02Z")
	t2 := mtime("2015-03-24T02:02:02Z")

	bs := newBits()
	bs.writeBits(start, 64)
	bs.writeBits(t1 - start, 14)
	dod := int64((t2 - t1) - (t1 - start))
	bs.writeBits(uint64(0b10), 2)
	bs.writeBits(uint64(dod), 7)
	// b8       [0 0 0 0 85 16 197 32]
	// expected [0 0 0 0 85 16 197 32 0 250 252]
	log.Printf("%v", bs.buf)
	r := newReader(bs.buf)
	var b8 [8]byte
	binary.BigEndian.PutUint64(b8[:], start)
	for i := 0; i < 8; i++ {
		b, err := r.readByte()
		assert.Nil(t, err)
		assert.Equal(t, b8[i], b)
	}
	delta, _ := r.readBits(14)
	assert.Equal(t, uint64(0b1111_10), delta)
	assert.Equal(t, t1-start, delta)
	dict, _ := r.readBits(2)
	assert.Equal(t, uint64(0b10), dict)
	log.Printf("%v", r.buf)
	v, _ := r.readBits(7)
	log.Printf("%d %v", v, r.buf)
	v2 := uint2int(v, 7)
	t.Logf("%d %d %d", v, int64(v), v2)
	assert.Equal(t, int64(-2), v2)
}

func TestReadNegative2(t *testing.T)  {
	bs := newBits()
	bs.writeBits(uint64(0b10), 2)
	three := uint64(3)
	five := uint64(5)
	a := int64(three - five)
	bs.writeBits(uint64(a), 7)
	t.Logf("%v", bs.buf)
	br := newReader(bs.buf)
	br.readBits(2)
	v, _ := br.readBits(7)
	v2 := uint2int(v, 7)
	t.Logf("%d %d %d", v, int64(v), v2)
	assert.Equal(t, int64(-2), v2)
}

func TestReadNegative3(t *testing.T)  {
	three := uint64(3)
	five := uint64(5)
	dod := int64(three - five)
	bs := newBits()
	bs.writeBits(62, 14)
	bs.writeBits(uint64(0b10), 2)
	bs.writeBits(uint64(dod), 7)
	// [0 250 252]
	t.Logf("%v", bs.buf)
	br := newReader(bs.buf)
	v, _ := br.readBits(14)
	assert.Equal(t, uint64(62), v)
	assert.Equal(t, uint8(2), br.remain)

	v, _ = br.readBits(2)
	assert.Equal(t, uint64(0b10), v)
	assert.Equal(t, uint8(0), br.remain)

	v, err := br.readBits(7)
	assert.Nil(t, err)
	assert.Equal(t, uint8(1), br.remain)
	t.Logf("%d", v)
	v2 := uint2int(v, 7)
	t.Logf("%d %d %d", v, int64(v), v2)
	assert.Equal(t, int64(-2), v2)
}

// given a RFC3339 string returns a unix epoch, panic if failed to convert
// https://github.com/golang/go/issues/9346
// The time.RFC3339 format is a case where the format string itself isn't a valid time. You can't have a Z and an offset in the time string, but the format string has both because the spec can contain either type of timezone specification.
//
// Both of these are valid RFC3339 times:
//
// "2015-09-15T14:00:12-00:00"
// "2015-09-15T14:00:13Z"
//
//And the time package needs to be able to parse them both using the same RFC3339 format string.
func mtime(s string) uint64 {
	tm, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return uint64(tm.Unix())
}
