package akumuli

import (
	"strings"
	"testing"

	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"
	asst "github.com/stretchr/testify/assert"
)

func TestEncoder_WritePointIntTagged(t *testing.T) {
	assert := asst.New(t)

	e := NewEncoder()
	p := pb.PointIntTagged{
		Name: "balancers.cpuload",
		Tags: []pb.Tag{
			{K: "host", V: "machine1"},
			{K: "region", V: "NW"},
		},
		Point: pb.PointInt{T: 1418224205000000000, V: 22},
	}
	p2 := pb.PointIntTagged{
		Name: "balancers.memusage",
		Tags: []pb.Tag{
			{K: "host", V: "machine1"},
			{K: "region", V: "NW"},
		},
		Point: pb.PointInt{T: 1418224205000000000, V: 23},
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
	assert.Equal(s, string(e.Bytes()))
	//t.Log(string(e.Bytes()))
}

func TestEncoder_WriteSeriesDoubleTagged(t *testing.T) {
	assert := asst.New(t)

	sd := &pb.SeriesDoubleTagged{
		Name: "archive_file_tracked",
		Tags: []pb.Tag{
			{K: "host", V: "server1"},
			{K: "data_center", V: "dc1"},
		},
		//1418224205000000000
		//1359788100000000000
		Points: []pb.PointDouble{
			{T: 1359788100000000000, V: 12.2},
			{T: 1359788200000000000, V: 13.3},
			{T: 1359788300000000000, V: 14.25},
		},
	}
	e := NewEncoder()
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
	assert.Equal(s, string(e.Bytes()))
}
