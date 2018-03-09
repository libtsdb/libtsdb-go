package genericw

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"
	"time"

	"github.com/dyweb/gommon/errors"
	"github.com/dyweb/gommon/requests"

	"github.com/libtsdb/libtsdb-go/libtsdb"
	"github.com/libtsdb/libtsdb-go/libtsdb/common"
	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"
	"github.com/libtsdb/libtsdb-go/libtsdb/util/bytesutil"
)

var _ libtsdb.TSDBClient = (*Client)(nil)
var _ libtsdb.WriteClient = (*Client)(nil)
var _ libtsdb.TracedHttpClient = (*Client)(nil)
var _ libtsdb.HttpClient = (*Client)(nil)

// Client is a generic HTTP based client for write, it is not go routine safe because encoder
// TODO: allow insecure, because we have https server with self signed certs, and HTTP/2 can only be used with https
type Client struct {
	enc     common.Encoder
	h       *http.Client
	baseReq *http.Request
	meta    libtsdb.Meta

	// flag for using http trace
	enableTrace bool

	// stat collected by client

	// single request
	// TODO: compressed
	proto string
	trace libtsdb.HttpTrace

	// accumulated counters
	bytesSend          uint64
	bytesSendSuccess   uint64
	intPointWritten    uint64
	doublePointWritten uint64
}

func New(meta libtsdb.Meta, encoder common.Encoder, req *http.Request) *Client {
	return &Client{
		enc:     encoder,
		h:       requests.NewDefaultClient(),
		baseReq: req,
		meta:    meta,
	}
}

func (c *Client) EnableHttpTrace() {
	c.enableTrace = true
}

func (c *Client) DisableHttpTrace() {
	c.enableTrace = false
}

func (c *Client) Meta() libtsdb.Meta {
	return c.meta
}

func (c *Client) SetHttpClient(h *http.Client) {
	c.h = h
}

// WriteIntPoint only writes to encoder, but does not flush it
func (c *Client) WriteIntPoint(p *pb.PointIntTagged) {
	c.intPointWritten += 1
	c.enc.WritePointIntTagged(p)
}

// WriteDoublePoint only writes to encoder, but does not flush it
func (c *Client) WriteDoublePoint(p *pb.PointDoubleTagged) {
	c.doublePointWritten += 1
	c.enc.WritePointDoubleTagged(p)
}

func (c *Client) WriteSeriesIntTagged(p *pb.SeriesIntTagged) {
	c.intPointWritten += uint64(len(p.Points))
	c.enc.WriteSeriesIntTagged(p)
}

func (c *Client) WriteSeriesDoubleTagged(p *pb.SeriesDoubleTagged) {
	c.doublePointWritten += uint64(len(p.Points))
	c.enc.WriteSeriesDoubleTagged(p)
}

// Flush sends encoded data to server and reset encoder
func (c *Client) Flush() error {
	return c.send()
}

func (c *Client) Trace() libtsdb.HttpTrace {
	return c.trace
}

func (c *Client) HttpStatusCode() int {
	return c.trace.StatusCode
}

func (c *Client) send() error {
	// TODO: real bytes send also include header etc, which we didn't take into account of bytes send
	payloadSize := uint64(c.enc.Len())
	c.bytesSend += payloadSize

	req := &http.Request{}
	*req = *c.baseReq
	b := c.enc.Bytes()
	req.Body = bytesutil.ReadCloser(b)

	// trace based on https://github.com/rakyll/hey/blob/master/requester/requester.go#L141
	trace := &c.trace
	trace.Start = time.Now().UnixNano()
	if c.enableTrace {
		tracer := &httptrace.ClientTrace{
			DNSStart: func(info httptrace.DNSStartInfo) {
				trace.DNSStart = time.Now().UnixNano()
			},
			DNSDone: func(info httptrace.DNSDoneInfo) {
				trace.DNSDone = time.Now().UnixNano()
			},
			// TODO: can we just ignore ConnectStart and ConnectDone?
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
	res, err := c.h.Do(req)
	c.enc.Reset()
	if err != nil {
		return errors.Wrap(err, "error send http request")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.Wrap(err, "can't read response body")
	}
	trace.StatusCode = res.StatusCode
	trace.ResDone = time.Now().UnixNano()
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent {
		// FIXED: log due to https://github.com/xephonhq/xephon-b/issues/36
		//log.Debugf("%d %s", res.StatusCode, string(b))
		return errors.New(string(body))
	}
	c.proto = res.Proto
	c.bytesSendSuccess += payloadSize
	return nil
}
