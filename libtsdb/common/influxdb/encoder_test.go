package influxdb

import (
	"testing"

	asst "github.com/stretchr/testify/assert"

	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"
)

func TestEncoder_WritePointIntTagged(t *testing.T) {
	assert := asst.New(t)
	p := &pb.PointIntTagged{
		Name:  "temperature",
		Point: pb.PointInt{T: int64(1434055562000000035), V: 35},
		Tags: []pb.Tag{
			{K: "machine", V: "unit42"},
			{K: "type", V: "assembly"},
		},
	}
	enc := NewEncoder()
	enc.WritePointIntTagged(p)
	assert.Equal("temperature,machine=unit42,type=assembly v=35 1434055562000000035", string(enc.Bytes()))

	c := enc.Cap()
	enc.Reset()
	assert.Equal(0, enc.Len())
	assert.Equal(c, enc.Cap())
	enc.WritePointIntTagged(p)
	assert.Equal("temperature,machine=unit42,type=assembly v=35 1434055562000000035", string(enc.Bytes()))
}

func TestEncoder_WritePointDoubleTagged(t *testing.T) {
	assert := asst.New(t)
	p := &pb.PointDoubleTagged{
		Name:  "temperature",
		Point: pb.PointDouble{T: int64(1434055562000000035), V: 35.132},
		Tags: []pb.Tag{
			{K: "machine", V: "unit42"},
			{K: "type", V: "assembly"},
		},
	}
	enc := NewEncoder()
	enc.WritePointDoubleTagged(p)
	assert.Equal("temperature,machine=unit42,type=assembly v=35.132 1434055562000000035", string(enc.Bytes()))
}
