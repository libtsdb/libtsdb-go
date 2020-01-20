package client_test

import (
	"testing"

	"github.com/libtsdb/libtsdb-go/database/kairosdb/client"
	"github.com/libtsdb/libtsdb-go/tspb"
	"github.com/stretchr/testify/assert"
)

func TestJsonEncoder_WritePointIntTagged(t *testing.T) {

	p := &tspb.PointIntTagged{
		Name:  "archive_file_search",
		Point: tspb.PointInt{Time: int64(1359786400000), Value: 321},
		Tags: []tspb.Tag{
			{Key: "host", Value: "server2"},
			{Key: "region", Value: "en-us"},
		},
	}
	enc := client.NewKairosDBJSONEncoder()
	enc.WritePointIntTagged(p)
	assert.Equal(t, `[{"name":"archive_file_search","timestamp":1359786400000,"value":321,"tags":{"host":"server2","region":"en-us"}}]`, string(enc.Bytes()))
	assert.Equal(t, `[{"name":"archive_file_search","timestamp":1359786400000,"value":321,"tags":{"host":"server2","region":"en-us"}}]`, string(enc.Bytes()))
	// check reset, found by https://github.com/xephonhq/xephon-b/issues/36
	enc.Reset()
	enc.WritePointIntTagged(p)
	assert.Equal(t, `[{"name":"archive_file_search","timestamp":1359786400000,"value":321,"tags":{"host":"server2","region":"en-us"}}]`, string(enc.Bytes()))
}

func TestJsonEncoder_WritePointDoubleTagged(t *testing.T) {

	p := &tspb.PointDoubleTagged{
		Name:  "cpu_idle",
		Point: tspb.PointDouble{Time: int64(1359786400000), Value: 23.2},
		Tags: []tspb.Tag{
			{Key: "host", Value: "server2"},
			{Key: "region", Value: "en-us"},
		},
	}
	enc := client.NewKairosDBJSONEncoder()
	enc.WritePointDoubleTagged(p)
	assert.Equal(t, `[{"name":"cpu_idle","timestamp":1359786400000,"value":23.2,"tags":{"host":"server2","region":"en-us"}}]`, string(enc.Bytes()))
}

func TestJsonEncoder_WriteSeriesIntTagged(t *testing.T) {
	s := &tspb.SeriesIntTagged{
		Name: "archive_file_tracked",
		Tags: []tspb.Tag{
			{Key: "host", Value: "server1"},
			{Key: "data_center", Value: "dc1"},
		},
		Points: []tspb.PointInt{
			{Time: 1359788100000, Value: 12},
			{Time: 1359788200000, Value: 13},
			{Time: 1359788300000, Value: 14},
		},
	}
	enc := client.NewKairosDBJSONEncoder()
	enc.WriteSeriesIntTagged(s)
	assert.Equal(t, `[{"name":"archive_file_tracked","datapoints":[[1359788100000,12],[1359788200000,13],[1359788300000,14]],"tags":{"host":"server1","data_center":"dc1"}}]`, string(enc.Bytes()))
}

func TestJsonEncoder_WriteSeriesDoubleTagged(t *testing.T) {
	s := &tspb.SeriesDoubleTagged{
		Name: "archive_file_tracked",
		Tags: []tspb.Tag{
			{Key: "host", Value: "server1"},
			{Key: "data_center", Value: "dc1"},
		},
		Points: []tspb.PointDouble{
			{Time: 1359788100000, Value: 12.2},
			{Time: 1359788200000, Value: 13.3},
			{Time: 1359788300000, Value: 14.25},
		},
	}
	enc := client.NewKairosDBJSONEncoder()
	enc.WriteSeriesDoubleTagged(s)
	assert.Equal(t, `[{"name":"archive_file_tracked","datapoints":[[1359788100000,12.2],[1359788200000,13.3],[1359788300000,14.25]],"tags":{"host":"server1","data_center":"dc1"}}]`, string(enc.Bytes()))
}

func TestTelnetEncoder_WritePointIntTagged(t *testing.T) {
	p := &tspb.PointIntTagged{
		Name:  "archive_file_search",
		Point: tspb.PointInt{Time: int64(1359786400000), Value: 321},
		Tags: []tspb.Tag{
			{Key: "host", Value: "server2"},
			{Key: "region", Value: "en-us"},
		},
	}
	enc := client.NewKaiorsDBTelnetEncoder()
	enc.WritePointIntTagged(p)
	s := `putm archive_file_search 1359786400000 321 host=server2 region=en-us
`
	assert.Equal(t, s, string(enc.Bytes()))
}

func TestTelnetEncoder_WriteSeriesDoubleTagged(t *testing.T) {

	sdt := &tspb.SeriesDoubleTagged{
		Name: "archive_file_tracked",
		Tags: []tspb.Tag{
			{Key: "host", Value: "server1"},
			{Key: "data_center", Value: "dc1"},
		},
		Points: []tspb.PointDouble{
			{Time: 1359788100000, Value: 12.2},
			{Time: 1359788200000, Value: 13.3},
			{Time: 1359788300000, Value: 14.25},
		},
	}
	enc := client.NewKaiorsDBTelnetEncoder()
	enc.WriteSeriesDoubleTagged(sdt)
	s := `putm archive_file_tracked 1359788100000 12.2 host=server1 data_center=dc1
putm archive_file_tracked 1359788200000 13.3 host=server1 data_center=dc1
putm archive_file_tracked 1359788300000 14.25 host=server1 data_center=dc1
`
	assert.Equal(t, s, string(enc.Bytes()))
}
