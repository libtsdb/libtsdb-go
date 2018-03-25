package xephonkgrpc

import (
	"testing"

	asst "github.com/stretchr/testify/assert"

	"github.com/libtsdb/libtsdb-go/libtsdb/config"
	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"
)

func TestClient_WriteIntPoint(t *testing.T) {
	t.Skipf("requires Xephon-K server running")

	assert := asst.New(t)

	c, err := New(*config.NewXephonkClientConfig())
	assert.Nil(err)
	c.WriteIntPoint(&pb.PointIntTagged{
		Name:  "temperature",
		Point: pb.PointInt{T: int64(1434055562000000035), V: 35},
		Tags: []pb.Tag{
			{K: "machine", V: "unit42"},
			{K: "type", V: "assembly"},
		},
	})
	err = c.Flush()
	// FIXME: xk currently does not have any storage
	t.Log(err)
}
