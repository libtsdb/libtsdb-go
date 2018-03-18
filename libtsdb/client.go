package libtsdb

import (
	"net/http"

	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"
)

// TSBClient returns meta of the database including its protocol, data type support
type TSDBClient interface {
	Meta() Meta
	Close() error
}

type WriteClient interface {
	TSDBClient
	WriteIntPoint(*pb.PointIntTagged)
	WriteDoublePoint(*pb.PointDoubleTagged)
	WriteSeriesIntTagged(p *pb.SeriesIntTagged)
	WriteSeriesDoubleTagged(p *pb.SeriesDoubleTagged)
	Flush() error
}

type HttpClient interface {
	SetHttpClient(client *http.Client)
	AllowInsecure()
	HttpStatusCode() int
}

type HttpWriteClient interface {
	WriteClient
	HttpClient
}

type TracedClient interface {
	Trace() Trace
}

type TracedHttpClient interface {
	TracedClient
	EnableHttpTrace()
	DisableHttpTrace()
	HttpTrace() HttpTrace
}

type TracedTcpClient interface {
	TracedClient
	TcpTrace() TcpTrace
}

// TODO: figure out the interface for read request...
type ReadClient interface {
	TSDBClient
	CreateDatabase(db string) error
}
