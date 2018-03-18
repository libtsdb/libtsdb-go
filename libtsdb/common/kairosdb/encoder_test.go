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
	assert.Equal(`[{"name":"archive_file_search","timestamp":1359786400000,"value":321,"tags":{"host":"server2","region":"en-us"}}]`, string(enc.Bytes()))
	assert.Equal(`[{"name":"archive_file_search","timestamp":1359786400000,"value":321,"tags":{"host":"server2","region":"en-us"}}]`, string(enc.Bytes()))
	// check reset, found by https://github.com/xephonhq/xephon-b/issues/36
	enc.Reset()
	enc.WritePointIntTagged(p)
	assert.Equal(`[{"name":"archive_file_search","timestamp":1359786400000,"value":321,"tags":{"host":"server2","region":"en-us"}}]`, string(enc.Bytes()))
}

func TestJsonEncoder_WritePointDoubleTagged(t *testing.T) {
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
	enc.WritePointDoubleTagged(p)
	assert.Equal(`[{"name":"cpu_idle","timestamp":1359786400000,"value":23.2,"tags":{"host":"server2","region":"en-us"}}]`, string(enc.Bytes()))
}

func TestJsonEncoder_WriteSeriesIntTagged(t *testing.T) {
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
	enc := NewJsonEncoder()
	enc.WriteSeriesIntTagged(s)
	assert.Equal(`[{"name":"archive_file_tracked","datapoints":[[1359788100000,12],[1359788200000,13],[1359788300000,14]],"tags":{"host":"server1","data_center":"dc1"}}]`, string(enc.Bytes()))
}

func TestJsonEncoder_WriteSeriesDoubleTagged(t *testing.T) {
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
	enc := NewJsonEncoder()
	enc.WriteSeriesDoubleTagged(s)
	assert.Equal(`[{"name":"archive_file_tracked","datapoints":[[1359788100000,12.2],[1359788200000,13.3],[1359788300000,14.25]],"tags":{"host":"server1","data_center":"dc1"}}]`, string(enc.Bytes()))
}

func TestTelnetEncoder_WritePointIntTagged(t *testing.T) {
	assert := asst.New(t)

	p := &pb.PointIntTagged{
		Name:  "archive_file_search",
		Point: pb.PointInt{T: int64(1359786400000), V: 321},
		Tags: []pb.Tag{
			{K: "host", V: "server2"},
			{K: "region", V: "en-us"},
		},
	}
	enc := NewTelnetEncoder()
	enc.WritePointIntTagged(p)
	s := `putm archive_file_search 1359786400000 321 host=server2 region=en-us
`
	assert.Equal(s, string(enc.Bytes()))
}

func TestTelnetEncoder_WriteSeriesDoubleTagged(t *testing.T) {
	assert := asst.New(t)

	sdt := &pb.SeriesDoubleTagged{
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
	enc := NewTelnetEncoder()
	enc.WriteSeriesDoubleTagged(sdt)
	s := `putm archive_file_tracked 1359788100000 12.2 host=server1 data_center=dc1
putm archive_file_tracked 1359788200000 13.3 host=server1 data_center=dc1
putm archive_file_tracked 1359788300000 14.25 host=server1 data_center=dc1
`
	assert.Equal(s, string(enc.Bytes()))
}
