package influxdbw

import (
	"testing"

	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"
	asst "github.com/stretchr/testify/assert"
)

// TODO: add flag to toggle test base on environ variable ... maybe testutil to gommon, travis etc.
func TestClient_WriteIntPoint(t *testing.T) {
	t.Skip("requires influxdb running")

	assert := asst.New(t)
	c, err := New(Config{
		Addr:     "http://localhost:8086",
		Database: "libtsdbtest",
	})
	assert.Nil(err)
	// TODO: util for point generator
	c.WriteIntPoint(&pb.PointIntTagged{
		Name:  "temperaturei",
		Point: pb.PointInt{T: int64(1434055562000000035), V: 35},
		Tags: []pb.Tag{
			{K: "machine", V: "unit42"},
			{K: "type", V: "assembly"},
		},
	})
	err = c.Flush()
	assert.Nil(err)
}

func TestClient_WriteDoublePoint(t *testing.T) {
	t.Skip("requires influxdb running")

	assert := asst.New(t)
	c, err := New(Config{
		Addr:     "http://localhost:8086",
		Database: "libtsdbtest",
	})
	assert.Nil(err)
	// TODO: influxdb even allow different type in a same series?
	c.WriteDoublePoint(&pb.PointDoubleTagged{
		Name:  "temperatured",
		Point: pb.PointDouble{T: int64(1434055562000000036), V: 35.132},
		Tags: []pb.Tag{
			{K: "machine", V: "unit42"},
			{K: "type", V: "assembly"},
		},
	})
	err = c.Flush()
	assert.Nil(err)
}
