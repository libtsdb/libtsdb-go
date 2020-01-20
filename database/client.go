package database

import (
	"net/http"

	"github.com/libtsdb/libtsdb-go/protocol"
	"github.com/libtsdb/libtsdb-go/tspb"
)

// TSBClient returns meta of the database including its protocol, data type support
type TSDBClient interface {
	Meta() Meta
	Close() error
}

type WriteClient interface {
	TSDBClient
	WriteIntPoint(*tspb.PointIntTagged)
	WriteDoublePoint(*tspb.PointDoubleTagged)
	WriteSeriesIntTagged(p *tspb.SeriesIntTagged)
	WriteSeriesDoubleTagged(p *tspb.SeriesDoubleTagged)
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
	Trace() protocol.Trace
}

type TracedWriteClient interface {
	TracedClient
	WriteClient
}

type TracedHttpClient interface {
	TracedClient
	EnableHttpTrace()
	DisableHttpTrace()
	HttpTrace() protocol.HttpTrace
}

type TracedTcpClient interface {
	TracedClient
	TcpTrace() protocol.TcpTrace
}
