package kairosdbw

import (
	"testing"

	"github.com/libtsdb/libtsdb-go/libtsdb/config"
	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"
	asst "github.com/stretchr/testify/assert"
)

func TestClient_WriteIntPoint(t *testing.T) {
	t.Skip("require kairosdb running")

	assert := asst.New(t)
	c, err := New(*config.NewKaiorsdbClientConfig())
	assert.Nil(err)
	c.WriteIntPoint(&pb.PointIntTagged{
		Name:  "archive_file_search",
		Point: pb.PointInt{T: int64(15198719140000), V: 321},
		Tags: []pb.Tag{
			{K: "host", V: "server2"},
			{K: "region", V: "en-us"},
		},
	})
	err = c.Flush()
	assert.Nil(err)
	c.WriteIntPoint(&pb.PointIntTagged{
		Name:  "archive_file_search",
		Point: pb.PointInt{T: int64(15198719150000), V: 322},
		Tags: []pb.Tag{
			{K: "host", V: "server2"},
			{K: "region", V: "en-us"},
		},
	})
	err = c.Flush()
	assert.Nil(err)

	// TODO: test server runs using a mock server to avoid issues like https://github.com/xephonhq/xephon-b/issues/36
}
