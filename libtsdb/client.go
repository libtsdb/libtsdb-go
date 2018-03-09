package libtsdb

import (
	"net/http"

	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"
)

// TODO: need to return status code etc.
type WriteClient interface {
	Meta() Meta
	WriteIntPoint(*pb.PointIntTagged)
	WriteDoublePoint(*pb.PointDoubleTagged)
	Flush() error
}

type TracedHttpClient interface {
	EnableHttpTrace()
	DisableHttpTrace()
	Trace() HttpTrace
}

type HttpClient interface {
	SetHttpClient(client *http.Client)
	HttpStatusCode() int
}

type HttpWriteClient interface {
	WriteClient
	HttpClient
}

// TODO: figure out the interface for read request...
type ReadClient interface {
	CreateDatabase(db string) error
	Meta() Meta
}
