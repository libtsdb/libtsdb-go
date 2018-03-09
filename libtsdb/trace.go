package libtsdb

// TODO: it might be more efficient to use unix timestamp
type HttpTrace struct {
	StatusCode int
	Start      int64
	DNSStart   int64
	DNSDone    int64
	GetConn    int64
	GotConn    int64
	Reused     bool
	TLSStart   int64
	TLSStop    int64
	ReqStart   int64
	ReqDone    int64
	ResStart   int64
	ResDone    int64
}
