package xephonk

import (
	"testing"

	asst "github.com/stretchr/testify/assert"

	"github.com/libtsdb/libtsdb-go/libtsdb/config"
	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"
)

func TestClient_WriteIntPoint(t *testing.T) {
	assert := asst.New(t)

	c, err := New(*config.NewXephonkClientConfig())
	assert.Nil(err)
	c.WriteIntPoint(&pb.PointIntTagged{
		Name:  "temperaturei",
		Point: pb.PointInt{T: int64(1434055562000000035), V: 35},
		Tags: []pb.Tag{
			{K: "machine", V: "unit42"},
			{K: "type", V: "assembly"},
		},
	})
	err = c.Flush()
	// FIXME: xk currently is not implemented
	t.Log(err)
}
