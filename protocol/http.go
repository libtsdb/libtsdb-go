package protocol

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"
	"time"

	"github.com/dyweb/gommon/errors"
	"github.com/dyweb/gommon/util/httputil"
	"github.com/libtsdb/libtsdb-go/database"
	"github.com/libtsdb/libtsdb-go/tspb"
	"github.com/libtsdb/libtsdb-go/util/bytesutil"
)

// HttpClient is a generic HTTP based client for write, it is not go routine safe because encoder
type HttpClient struct {
	// tsdb
	enc  Encoder
	meta database.Meta

	// http
	h           *http.Client
	insecure    bool
	baseReq     *http.Request
	enableTrace bool // use net/http/httptrace

	// stat for single request
	// TODO: deal with gzip
	proto string
	trace HttpTrace
	// TODO: why keep previous trace
	prevTrace HttpTrace

	// stat for accumulated counters
	// NOTE: we maintain counter in generic clients so encoder don't need to worry about it
	totalPayloadSize        int
	totalRawSize            int
	totalRawMetaSize        int
	totalIntPointWritten    int
	totalDoublePointWritten int
	// TODO: can't count unique series written unless we hash series
}

func NewHttp(meta database.Meta, encoder Encoder, req *http.Request) *HttpClient {
	return &HttpClient{
		enc:     encoder,
		h:       httputil.NewUnPooledClient(),
		baseReq: req,
		meta:    meta,
	}
}

func (c *HttpClient) EnableHttpTrace() {
	c.enableTrace = true
}

func (c *HttpClient) DisableHttpTrace() {
	c.enableTrace = false
}

func (c *HttpClient) AllowInsecure() {
	if c.h == nil {
		return
	}
	c.insecure = true
	if t, ok := c.h.Transport.(*http.Transport); ok {
		t.TLSClientConfig.InsecureSkipVerify = true
	}
}

func (c *HttpClient) Meta() database.Meta {
	return c.meta
}

func (c *HttpClient) Close() error {
	// http client doesn't not have methods for closing it ...
	return nil
}

func (c *HttpClient) SetHttpClient(h *http.Client) {
	c.h = h
	// TODO: maybe we should not set insecure because the user can set it by themselve since they are already
	// setting the http client directly ...
	if c.insecure {
		if t, ok := c.h.Transport.(*http.Transport); ok {
			t.TLSClientConfig.InsecureSkipVerify = true
		}
	}
}

// WriteIntPoint only writes to encoder, but does not flush it
func (c *HttpClient) WriteIntPoint(p *tspb.PointIntTagged) {
	c.trace.Points += 1
	c.trace.RawSize += p.RawSize()
	c.trace.RawMetaSize += p.RawMetaSize()

	c.totalIntPointWritten += 1
	c.totalRawSize += p.RawSize()
	c.totalRawMetaSize += p.RawMetaSize()

	c.enc.WritePointIntTagged(p)
}

// WriteDoublePoint only writes to encoder, but does not flush it
func (c *HttpClient) WriteDoublePoint(p *tspb.PointDoubleTagged) {
	c.trace.Points += 1
	c.trace.RawSize += p.RawSize()
	c.trace.RawMetaSize += p.RawMetaSize()

	c.totalDoublePointWritten += 1
	c.totalRawSize += p.RawSize()
	c.totalRawMetaSize += p.RawMetaSize()

	c.enc.WritePointDoubleTagged(p)
}

func (c *HttpClient) WriteSeriesIntTagged(p *tspb.SeriesIntTagged) {
	c.trace.Points += len(p.Points)
	c.trace.RawSize += p.RawSize()
	c.trace.RawMetaSize += p.RawMetaSize()

	c.totalIntPointWritten += len(p.Points)
	c.totalRawSize += p.RawSize()
	c.totalRawMetaSize += p.RawMetaSize()

	c.enc.WriteSeriesIntTagged(p)
}

func (c *HttpClient) WriteSeriesDoubleTagged(p *tspb.SeriesDoubleTagged) {
	c.trace.Points += len(p.Points)
	c.trace.RawSize += p.RawSize()
	c.trace.RawMetaSize += p.RawMetaSize()

	c.totalDoublePointWritten += len(p.Points)
	c.totalRawSize += p.RawSize()
	c.totalRawMetaSize += p.RawMetaSize()

	c.enc.WriteSeriesDoubleTagged(p)
}

// Flush sends encoded data to server and reset encoder
func (c *HttpClient) Flush() error {
	return c.send()
}

func (c *HttpClient) Trace() Trace {
	// make a copy, otherwise when the trace is used, the pointer might be pointing to a changed trace
	cp := c.prevTrace
	return &cp
}

func (c *HttpClient) HttpTrace() HttpTrace {
	return c.prevTrace
}

func (c *HttpClient) HttpStatusCode() int {
	return c.prevTrace.StatusCode
}

func (c *HttpClient) send() error {
	// TODO: real bytes send also include http header etc, which we didn't take into account of bytes send
	c.totalPayloadSize += c.enc.Len()
	c.trace.PayloadSize = c.enc.Len()

	req := &http.Request{}
	*req = *c.baseReq
	b := c.enc.Bytes()
	req.Body = bytesutil.ReadCloser(b)

	// trace based on https://github.com/rakyll/hey/blob/master/requester/requester.go#L141
	trace := &c.trace
	if c.enableTrace {
		tracer := &httptrace.ClientTrace{
			DNSStart: func(info httptrace.DNSStartInfo) {
				trace.DNSStart = time.Now().UnixNano()
			},
			DNSDone: func(info httptrace.DNSDoneInfo) {
				trace.DNSDone = time.Now().UnixNano()
			},
			// TODO: is it ok to ignore ConnectStart and ConnectDone?
			GetConn: func(hostPort string) {
				trace.GetConn = time.Now().UnixNano()
			},
			GotConn: func(info httptrace.GotConnInfo) {
				now := time.Now().UnixNano()
				trace.Reused = info.Reused
				trace.GotConn = now
				trace.ReqStart = now
			},
			// TODO: only tls handshake when new connection is established?
			TLSHandshakeStart: func() {
				trace.TLSStart = time.Now().UnixNano()
			},
			TLSHandshakeDone: func(state tls.ConnectionState, e error) {
				trace.TLSStop = time.Now().UnixNano()
			},
			WroteRequest: func(info httptrace.WroteRequestInfo) {
				trace.ReqDone = time.Now().UnixNano()
			},
			GotFirstResponseByte: func() {
				trace.ResStart = time.Now().UnixNano()
			},
		}
		req = req.WithContext(httptrace.WithClientTrace(req.Context(), tracer))
	}

	trace.Start = time.Now().UnixNano()
	res, err := c.h.Do(req)
	trace.End = time.Now().UnixNano()
	// reset
	c.enc.Reset()
	c.prevTrace = c.trace
	c.trace.Reset()
	if err != nil {
		c.prevTrace.Error = true
		c.prevTrace.ErrorMessage = err.Error()
		return errors.Wrap(err, "error send http request")
	}
	// read response
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.prevTrace.Error = true
		c.prevTrace.ErrorMessage = err.Error()
		return errors.Wrap(err, "can't read response body")
	}
	c.prevTrace.StatusCode = res.StatusCode
	c.prevTrace.ResDone = time.Now().UnixNano()
	c.prevTrace.End = c.prevTrace.ResDone // TODO: might need two types of end time, on for finished write, one for finished read
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent {
		// FIXED: log due to https://github.com/xephonhq/xephon-b/issues/36
		//log.Debugf("%d %s", res.StatusCode, string(b))
		c.prevTrace.Error = true
		c.prevTrace.ErrorMessage = string(body)
		return errors.New(string(body))
	}
	c.proto = res.Proto
	return nil
}
