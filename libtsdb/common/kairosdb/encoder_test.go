package kairosdb

import (
	"testing"

	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"
	asst "github.com/stretchr/testify/assert"
)

func TestJsonEncoder_WritePointIntTagged(t *testing.T) {
	assert := asst.New(t)

	p := &pb.PointIntTagged{
		Name:  "archive_file_search",
		Point: pb.PointInt{T: int64(1359786400000), V: 321},
		Tags: []pb.Tag{
			{K: "host", V: "server2"},
			{K: "region", V: "en-us"},
		},
	}
	enc := NewJsonEncoder()
	enc.WritePointIntTagged(p)
	assert.Equal(`{"name":"archive_file_search","timestamp":1359786400000,"value":321,"tags":{"host":"server2","region":"en-us"}}`, string(enc.Bytes()))
}

func TestJsonEncoder_WriteDoublePointTagged(t *testing.T) {
	assert := asst.New(t)

	p := &pb.PointDoubleTagged{
		Name:  "cpu_idle",
		Point: pb.PointDouble{T: int64(1359786400000), V: 23.2},
		Tags: []pb.Tag{
			{K: "host", V: "server2"},
			{K: "region", V: "en-us"},
		},
	}
	enc := NewJsonEncoder()
	enc.WriteDoublePointTagged(p)
	assert.Equal(`{"name":"cpu_idle","timestamp":1359786400000,"value":23.2,"tags":{"host":"server2","region":"en-us"}}`, string(enc.Bytes()))
}
