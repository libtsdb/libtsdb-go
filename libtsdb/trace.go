package libtsdb

import "time"

// TODO: it might be more efficient to use unix timestamp
type HttpTrace struct {
	StatusCode int
	Start      time.Time
	DNSStart   time.Time
	DNSDone    time.Time
	GetConn    time.Time
	GotConn    time.Time
	Reused     bool
	TLSStart   time.Time
	TLSStop    time.Time
	ReqStart   time.Time
	ReqDone    time.Time
	ResStart   time.Time
	ResDone    time.Time
}
