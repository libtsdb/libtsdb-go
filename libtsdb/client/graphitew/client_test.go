package graphitew

import (
	"testing"

	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"
	asst "github.com/stretchr/testify/assert"
)

func TestClient_WriteIntPoint(t *testing.T) {
	t.Skip("requires graphite running")

	assert := asst.New(t)
	c, err := New(Config{
		Addr: "localhost:2003",
	})
	assert.Nil(err)
	err = c.WriteIntPoint(&pb.PointIntTagged{
		Name: "temperature",
		// TODO: time precision ...
		Point: pb.PointInt{T: int64(1519266078), V: 35},
		Tags: []pb.Tag{
			{K: "machine", V: "unit42"},
			{K: "type", V: "assembly"},
		},
	})
	assert.Nil(err)
}

func TestClient_WriteDoublePoint(t *testing.T) {
	t.Skip("requires graphite running")

	assert := asst.New(t)
	c, err := New(Config{
		Addr: "localhost:2003",
	})
	assert.Nil(err)
	err = c.WriteDoublePoint(&pb.PointDoubleTagged{
		Name:  "temperature",
		Point: pb.PointDouble{T: int64(1519266079), V: 35.132},
		Tags: []pb.Tag{
			{K: "machine", V: "unit42"},
			{K: "type", V: "assembly"},
		},
	})
	assert.Nil(err)
}
