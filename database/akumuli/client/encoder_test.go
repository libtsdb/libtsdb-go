package client_test

import (
	"strings"
	"testing"

	"github.com/libtsdb/libtsdb-go/database/akumuli/client"
	"github.com/libtsdb/libtsdb-go/tspb"
	"github.com/stretchr/testify/assert"
)

func TestAkumuliEncoder_WritePointIntTagged(t *testing.T) {
	e := client.NewAkumuliEncoder()
	p := tspb.
		PointIntTagged{
		Name: "balancers.cpuload",
		Tags: []tspb.Tag{
			{Key: "host", Value: "machine1"},
			{Key: "region", Value: "NW"},
		},
		Point: tspb.PointInt{Time: 1418224205000000000, Value: 22},
	}
	p2 := tspb.
		PointIntTagged{
		Name: "balancers.memusage",
		Tags: []tspb.Tag{
			{Key: "host", Value: "machine1"},
			{Key: "region", Value: "NW"},
		},
		Point: tspb.PointInt{Time: 1418224205000000000, Value: 23},
	}
	e.WritePointIntTagged(&p)
	e.WritePointIntTagged(&p2)
	s := `+balancers.cpuload host=machine1 region=NW
:1418224205000000000
+22
+balancers.memusage host=machine1 region=NW
:1418224205000000000
+23
`
	s = strings.Replace(s, "\n", "\r\n", -1)
	assert.Equal(t, s, string(e.Bytes()))
}

func TestEncoder_WriteSeriesDoubleTagged(t *testing.T) {

	sd := &tspb.SeriesDoubleTagged{
		Name: "archive_file_tracked",
		Tags: []tspb.Tag{
			{Key: "host", Value: "server1"},
			{Key: "data_center", Value: "dc1"},
		},
		//1418224205000000000
		//1359788100000000000
		Points: []tspb.PointDouble{
			{Time: 1359788100000000000, Value: 12.2},
			{Time: 1359788200000000000, Value: 13.3},
			{Time: 1359788300000000000, Value: 14.25},
		},
	}
	e := client.NewAkumuliEncoder()
	e.WriteSeriesDoubleTagged(sd)
	s := `+archive_file_tracked host=server1 data_center=dc1
:1359788100000000000
+12.2
+archive_file_tracked host=server1 data_center=dc1
:1359788200000000000
+13.3
+archive_file_tracked host=server1 data_center=dc1
:1359788300000000000
+14.25
`
	s = strings.Replace(s, "\n", "\r\n", -1)
	assert.Equal(t, s, string(e.Bytes()))
}
