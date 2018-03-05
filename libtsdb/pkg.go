package libtsdb

import (
	"net/http"

	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"
)

// TODO: need to return status code etc.
type WriteClient interface {
	WriteIntPoint(*pb.PointIntTagged)
	WriteDoublePoint(*pb.PointDoubleTagged)
	Flush() error
}

type HttpClient interface {
	SetHttpClient(client *http.Client)
}

type ReadClient interface {
}
