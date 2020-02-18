# Prometheus Compression

Compression for data is in [tsdb/chunkenc](https://github.com/prometheus/prometheus/tree/master/tsdb/chunkenc).

## Encode

There is a `Appender` interface and two implementations, the second one is just a thread safe wrapper.
The main logic is in [xorAppender.Append](https://github.com/prometheus/prometheus/blob/0703dae7cc4fcb1e051ab5fec89c47530e78c75a/tsdb/chunkenc/xor.go#L149)
which contains the double delta logic for timestamp and deal with float64 in [xorAppender.writeVDelta](https://github.com/prometheus/prometheus/blob/0703dae7cc4fcb1e051ab5fec89c47530e78c75a/tsdb/chunkenc/xor.go#L206)

```go
// tsdb/chunenc/xor.go

type xorAppender struct {
	b *bstream

	t      int64
	v      float64
	tDelta uint64

	leading  uint8
	trailing uint8
}

func (a *xorAppender) Append(t int64, v float64) {
	var tDelta uint64
	num := binary.BigEndian.Uint16(a.b.bytes())

	if num == 0 {
		buf := make([]byte, binary.MaxVarintLen64)
		for _, b := range buf[:binary.PutVarint(buf, t)] {
			a.b.writeByte(b)
		}
		a.b.writeBits(math.Float64bits(v), 64)

	} else if num == 1 {
		tDelta = uint64(t - a.t)

		buf := make([]byte, binary.MaxVarintLen64)
		for _, b := range buf[:binary.PutUvarint(buf, tDelta)] {
			a.b.writeByte(b)
		}

		a.writeVDelta(v)

	} else {
		tDelta = uint64(t - a.t)
		dod := int64(tDelta - a.tDelta)

		// Gorilla has a max resolution of seconds, Prometheus milliseconds.
		// Thus we use higher value range steps with larger bit size.
		switch {
		case dod == 0:
			a.b.writeBit(zero)
		case bitRange(dod, 14):
			a.b.writeBits(0x02, 2) // '10'
			a.b.writeBits(uint64(dod), 14)
		case bitRange(dod, 17):
			a.b.writeBits(0x06, 3) // '110'
			a.b.writeBits(uint64(dod), 17)
		case bitRange(dod, 20):
			a.b.writeBits(0x0e, 4) // '1110'
			a.b.writeBits(uint64(dod), 20)
		default:
			a.b.writeBits(0x0f, 4) // '1111'
			a.b.writeBits(uint64(dod), 64)
		}

		a.writeVDelta(v)
	}

	a.t = t
	a.v = v
	binary.BigEndian.PutUint16(a.b.bytes(), num+1)
	a.tDelta = tDelta
}

func (a *xorAppender) writeVDelta(v float64) {
	vDelta := math.Float64bits(v) ^ math.Float64bits(a.v)

	if vDelta == 0 {
		a.b.writeBit(zero)
		return
	}
	a.b.writeBit(one)

	leading := uint8(bits.LeadingZeros64(vDelta))
	trailing := uint8(bits.TrailingZeros64(vDelta))

	// Clamp number of leading zeros to avoid overflow when encoding.
	if leading >= 32 {
		leading = 31
	}

	if a.leading != 0xff && leading >= a.leading && trailing >= a.trailing {
		a.b.writeBit(zero)
		a.b.writeBits(vDelta>>a.trailing, 64-int(a.leading)-int(a.trailing))
	} else {
		a.leading, a.trailing = leading, trailing

		a.b.writeBit(one)
		a.b.writeBits(uint64(leading), 5)

		// Note that if leading == trailing == 0, then sigbits == 64.  But that value doesn't actually fit into the 6 bits we have.
		// Luckily, we never need to encode 0 significant bits, since that would put us in the other case (vdelta == 0).
		// So instead we write out a 0 and adjust it back to 64 on unpacking.
		sigbits := 64 - leading - trailing
		a.b.writeBits(uint64(sigbits), 6)
		a.b.writeBits(vDelta>>trailing, int(sigbits))
	}
}
```