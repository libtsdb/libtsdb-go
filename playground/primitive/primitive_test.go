package primitive_test

import (
	"encoding/binary"
	"reflect"
	"testing"
	"unsafe"
)

// test primitive types

func TestEndianness(t *testing.T) {
	v := uint64(1024)
	var buf [8]byte
	// The implementation simply do right shift
	binary.BigEndian.PutUint64(buf[:], v)
	// [0 0 0 0 0 0 4 0], which is
	//  b[0] = byte(v >> 56)
	//	b[1] = byte(v >> 48)
	// ...
	//	b[6] = byte(v >> 8)
	//	b[7] = byte(v)
	t.Logf("%v", buf)

	// The implementation is also doing right shift ... just different order
	// _ = b[7] // early bounds check to guarantee safety of writes below
	//	b[0] = byte(v)
	//	b[1] = byte(v >> 8)
	//	b[2] = byte(v >> 16)
	// ...
	//	b[7] = byte(v >> 56)
	binary.LittleEndian.PutUint64(buf[:], v)
	t.Logf("%v", buf)

	// [0 4 0 0 0 0 0 0] it's little endian when using unsafe, ok ...
	t.Logf("%v", unsafeInt2Bytes(v))
}

// https://stackoverflow.com/a/17539687
func unsafeInt2Bytes(v uint64) []byte {
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&v)),
		Len:  8,
		Cap:  8,
	}
	return *(*[]byte)(unsafe.Pointer(&hdr))
}
