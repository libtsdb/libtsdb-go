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
	assert.Equal("temperature;machine=unit42;type=assembly 35 1519266078\n", string(enc.Bytes()))

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
	assert.Equal("temperature;machine=unit42;type=assembly 35.132 1519266078\n", string(enc.Bytes()))

	enc.Reset()
	assert.Equal(0, enc.Len())
}

func TestTextEncoder_WriteSeriesIntTagged(t *testing.T) {
	assert := asst.New(t)

	s := &pb.SeriesIntTagged{
		Name: "archive_file_tracked",
		Tags: []pb.Tag{
			{K: "host", V: "server1"},
			{K: "data_center", V: "dc1"},
		},
		Points: []pb.PointInt{
			{T: 1359788100000, V: 12},
			{T: 1359788200000, V: 13},
			{T: 1359788300000, V: 14},
		},
	}
	enc := NewTextEncoder()
	enc.WriteSeriesIntTagged(s)
	res := `archive_file_tracked;host=server1;data_center=dc1 12 1359788100000
archive_file_tracked;host=server1;data_center=dc1 13 1359788200000
archive_file_tracked;host=server1;data_center=dc1 14 1359788300000
`
	assert.Equal(res, string(enc.Bytes()))
}

func TestTextEncoder_WriteSeriesDoubleTagged(t *testing.T) {
	assert := asst.New(t)

	s := &pb.SeriesDoubleTagged{
		Name: "archive_file_tracked",
		Tags: []pb.Tag{
			{K: "host", V: "server1"},
			{K: "data_center", V: "dc1"},
		},
		Points: []pb.PointDouble{
			{T: 1359788100000, V: 12.2},
			{T: 1359788200000, V: 13.3},
			{T: 1359788300000, V: 14.25},
		},
	}
	enc := NewTextEncoder()
	enc.WriteSeriesDoubleTagged(s)
	res := `archive_file_tracked;host=server1;data_center=dc1 12.2 1359788100000
archive_file_tracked;host=server1;data_center=dc1 13.3 1359788200000
archive_file_tracked;host=server1;data_center=dc1 14.25 1359788300000
`
	assert.Equal(res, string(enc.Bytes()))
}
