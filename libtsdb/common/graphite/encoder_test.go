package graphite

import (
	"testing"

	asst "github.com/stretchr/testify/assert"

	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"
)

func TestTextEncoder_WritePointIntTagged(t *testing.T) {
	assert := asst.New(t)
	// TODO: test util
	p := &pb.PointIntTagged{
		Name: "temperature",
		// TODO: time precision ...
		Point: pb.PointInt{T: int64(1519266078), V: 35},
		Tags: []pb.Tag{
			{K: "machine", V: "unit42"},
			{K: "type", V: "assembly"},
		},
	}
	enc := NewTextEncoder()
	enc.WritePointIntTagged(p)
	assert.Equal("temperature;machine=unit42;type=assembly 35 1519266078", string(enc.Bytes()))

	enc.Reset()
	assert.Equal(0, enc.Len())
}

func TestTextEncoder_WritePointDoubleTagged(t *testing.T) {
	assert := asst.New(t)
	p := &pb.PointDoubleTagged{
		Name:  "temperature",
		Point: pb.PointDouble{T: int64(1519266078), V: 35.132},
		Tags: []pb.Tag{
			{K: "machine", V: "unit42"},
			{K: "type", V: "assembly"},
		},
	}
	enc := NewTextEncoder()
	enc.WritePointDoubleTagged(p)
	assert.Equal("temperature;machine=unit42;type=assembly 35.132 1519266078", string(enc.Bytes()))

	enc.Reset()
	assert.Equal(0, enc.Len())
}
