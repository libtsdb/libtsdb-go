package akumuliw

import (
	"testing"

	asst "github.com/stretchr/testify/assert"

	"github.com/libtsdb/libtsdb-go/libtsdb/config"
	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"
)

func TestClient_WriteIntPoint(t *testing.T) {
	t.Skip("requires akumuli running")

	assert := asst.New(t)

	c, err := New(*config.NewAkumuliClientConfig())
	assert.Nil(err)
	c.WriteIntPoint(&pb.PointIntTagged{
		Name: "balancers.cpuload",
		Tags: []pb.Tag{
			{K: "host", V: "machine1"},
			{K: "region", V: "NW"},
		},
		Point: pb.PointInt{T: 1418224205000000000, V: 22},
	})
	err = c.Flush()
	assert.Nil(err)
}
