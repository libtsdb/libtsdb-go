package client_test

import (
	"testing"

	"github.com/libtsdb/libtsdb-go/database/influxdb/client"
	"github.com/libtsdb/libtsdb-go/tspb"
	"github.com/stretchr/testify/assert"
)

func TestEncoder_WritePointIntTagged(t *testing.T) {
	// TODO: util for point generator
	p := &tspb.PointIntTagged{
		Name:  "temperature",
		Point: tspb.PointInt{Time: int64(1434055562000000035), Value: 35},
		Tags: []tspb.Tag{
			{Key: "machine", Value: "unit42"},
			{Key: "type", Value: "assembly"},
		},
	}
	enc := client.NewInfluxDBEncoder()
	enc.WritePointIntTagged(p)
	assert.Equal(t, "temperature,machine=unit42,type=assembly v=35 1434055562000000035\n", string(enc.Bytes()))

	c := enc.Cap()
	enc.Reset()
	assert.Equal(t, 0, enc.Len())
	assert.Equal(t, c, enc.Cap())
	enc.WritePointIntTagged(p)
	assert.Equal(t, "temperature,machine=unit42,type=assembly v=35 1434055562000000035\n", string(enc.Bytes()))
}

func TestEncoder_WritePointDoubleTagged(t *testing.T) {
	p := &tspb.PointDoubleTagged{
		Name:  "temperature",
		Point: tspb.PointDouble{Time: int64(1434055562000000035), Value: 35.132},
		Tags: []tspb.Tag{
			{Key: "machine", Value: "unit42"},
			{Key: "type", Value: "assembly"},
		},
	}
	enc := client.NewInfluxDBEncoder()
	enc.WritePointDoubleTagged(p)
	assert.Equal(t, "temperature,machine=unit42,type=assembly v=35.132 1434055562000000035\n", string(enc.Bytes()))
}

func TestEncoder_WriteSeriesIntTagged(t *testing.T) {

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
	enc := client.NewInfluxDBEncoder()
	enc.WriteSeriesIntTagged(s)
	res := `archive_file_tracked,host=server1,data_center=dc1 v=12 1359788100000
archive_file_tracked,host=server1,data_center=dc1 v=13 1359788200000
archive_file_tracked,host=server1,data_center=dc1 v=14 1359788300000
`
	assert.Equal(t, res, string(enc.Bytes()))
}

func TestEncoder_WriteSeriesDoubleTagged(t *testing.T) {
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
	enc := client.NewInfluxDBEncoder()
	enc.WriteSeriesDoubleTagged(s)
	res := `archive_file_tracked,host=server1,data_center=dc1 v=12.2 1359788100000
archive_file_tracked,host=server1,data_center=dc1 v=13.3 1359788200000
archive_file_tracked,host=server1,data_center=dc1 v=14.25 1359788300000
`
	assert.Equal(t, res, string(enc.Bytes()))
}
