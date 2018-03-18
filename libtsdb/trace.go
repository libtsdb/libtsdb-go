package libtsdb

// TODO: tcp client

type HttpTrace struct {
	// from response
	StatusCode   int
	Error        bool
	ErrorMessage string

	// TODO: RawSize, meta + points
	// TODO: MetaSize

	// from net/http/httptrace
	Start    int64
	DNSStart int64
	DNSDone  int64
	GetConn  int64
	GotConn  int64
	Reused   bool
	TLSStart int64
	TLSStop  int64
	ReqStart int64
	ReqDone  int64
	ResStart int64
	ResDone  int64
}
