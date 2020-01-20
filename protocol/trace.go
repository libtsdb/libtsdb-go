package protocol

var _ Trace = (*TcpTrace)(nil)
var _ Trace = (*HttpTrace)(nil)

type Trace interface {
	// GetError specifies if this result is an error
	GetError() bool
	// GetErrorMessage is empty if Error is false
	GetErrorMessage() string
	// GetCode is response code from server, normally http status code
	GetCode() int
	// GetStartTime is when the request is started
	GetStartTime() int64
	// GetEndTime is when the request is finished, response is drained, error or not
	GetEndTime() int64
	// GetPoints is number of points written in request
	GetPoints() int
	// GetPayloadSize is the size of the payload excluding header etc.
	GetPayloadSize() int
	// GetRawSize is the size in byte for meta and points written without serialization, see libtsdbpb sizer.go
	GetRawSize() int
	// GetRawMetaSize is the size in byte for meta data written without serialization, series name tags etc.
	GetRawMetaSize() int
}

type TcpTrace struct {
	// response TODO: we don't wait response cause many of them don't have response
	Code         int // TODO: any code like http status code in tcp ...
	Error        bool
	ErrorMessage string

	Points int
	// TODO: can't count unique series unless we hash

	// size
	PayloadSize int
	RawSize     int
	RawMetaSize int

	// time
	Start int64
	End   int64
}

func (t *TcpTrace) GetError() bool {
	return t.Error
}

func (t *TcpTrace) GetErrorMessage() string {
	return t.ErrorMessage
}

func (t *TcpTrace) GetCode() int {
	return t.Code
}

func (t *TcpTrace) GetStartTime() int64 {
	return t.Start
}

func (t *TcpTrace) GetEndTime() int64 {
	return t.End
}

func (t *TcpTrace) GetPoints() int {
	return t.Points
}

func (t *TcpTrace) GetPayloadSize() int {
	return t.PayloadSize
}

func (t *TcpTrace) GetRawSize() int {
	return t.RawSize
}

func (t *TcpTrace) GetRawMetaSize() int {
	return t.RawMetaSize
}

func (t *TcpTrace) Reset() {
	t.Code = 0
	t.Error = false
	t.ErrorMessage = ""

	t.Points = 0
	t.PayloadSize = 0
	t.RawSize = 0
	t.RawMetaSize = 0

	// time reset is not needed
}

type HttpTrace struct {
	// response
	StatusCode   int
	Error        bool
	ErrorMessage string

	Points int
	// TODO: can't count unique series unless we hash

	// size
	PayloadSize int
	RawSize     int
	RawMetaSize int

	// time
	Start int64
	End   int64
	// net/http/httptrace
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

func (t *HttpTrace) GetError() bool {
	return t.Error
}

func (t *HttpTrace) GetErrorMessage() string {
	return t.ErrorMessage
}

func (t *HttpTrace) GetCode() int {
	return t.StatusCode
}

func (t *HttpTrace) GetStartTime() int64 {
	return t.Start
}

func (t *HttpTrace) GetEndTime() int64 {
	return t.End
}

func (t *HttpTrace) GetPoints() int {
	return t.Points
}

func (t *HttpTrace) GetPayloadSize() int {
	return t.PayloadSize
}

func (t *HttpTrace) GetRawSize() int {
	return t.RawSize
}

func (t *HttpTrace) GetRawMetaSize() int {
	return t.RawMetaSize
}

func (t *HttpTrace) Reset() {
	t.StatusCode = 0
	t.Error = false
	t.ErrorMessage = ""

	t.Points = 0
	t.PayloadSize = 0
	t.RawSize = 0
	t.RawMetaSize = 0

	// time and httptrace reset is not needed
}
