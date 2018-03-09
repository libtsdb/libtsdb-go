package libtsdb

import "time"

// TODO: use in generic client
type HttpTrace struct {
	DnsDuration  time.Duration
	ConnDuration time.Duration
	TlsDuration  time.Duration
	ReqDuration  time.Duration
	ResDuration  time.Duration
}
