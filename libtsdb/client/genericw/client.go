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

var _ libtsdb.WriteClient = (*Client)(nil)
var _ libtsdb.TracedHttpClient = (*Client)(nil)
var _ libtsdb.HttpClient = (*Client)(nil)

// Client is a generic HTTP based client for write, it is not go routine safe because encoder
// TODO: allow insecure, because we have https server with self signed certs
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
	statusCode int
	proto      string
	trace      libtsdb.HttpTrace

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

// Flush sends encoded data to server and reset encoder
func (c *Client) Flush() error {
	return c.send()
}

func (c *Client) Trace() libtsdb.HttpTrace {
	return c.trace
}

func (c *Client) HttpStatusCode() int {
	return c.statusCode
}

func (c *Client) send() error {
	payloadSize := uint64(c.enc.Len())
	c.bytesSend += payloadSize

	// TODO: go support http client tracing, we can also use open tracing here ...
	// TODO: real bytes send also include header etc, which we didn't take into account of bytes send
	req := &http.Request{}
	*req = *c.baseReq
	b := c.enc.Bytes()
	req.Body = bytesutil.ReadCloser(b)
	// based on https://github.com/rakyll/hey/blob/master/requester/requester.go#L141
	var dnsStart, connStart, tlsStart, reqStart, resStart time.Time
	trace := &c.trace
	if c.enableTrace {
		tracer := &httptrace.ClientTrace{
			DNSStart: func(info httptrace.DNSStartInfo) {
				dnsStart = time.Now()
			},
			DNSDone: func(info httptrace.DNSDoneInfo) {
				trace.DnsDuration = time.Now().Sub(dnsStart)
			},
			// TODO: can we just ignore ConnectStart and ConnectDone?
			GetConn: func(hostPort string) {
				connStart = time.Now()
			},
			GotConn: func(info httptrace.GotConnInfo) {
				// TODO: info also contains Idle etc.
				now := time.Now()
				if info.Reused {
					trace.ConnDuration = 0
				} else {
					trace.ConnDuration = now.Sub(connStart)
				}
				reqStart = now
			},
			// TODO: only tls handshake when new connection is established?
			TLSHandshakeStart: func() {
				tlsStart = time.Now()
			},
			TLSHandshakeDone: func(state tls.ConnectionState, e error) {
				trace.TlsDuration = time.Now().Sub(tlsStart)
			},
			WroteRequest: func(info httptrace.WroteRequestInfo) {
				trace.ReqDuration = time.Now().Sub(reqStart)
			},
			GotFirstResponseByte: func() {
				resStart = time.Now()
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
	c.statusCode = res.StatusCode
	trace.ResDuration = time.Now().Sub(resStart)
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent {
		// TODO: might disable this since https://github.com/xephonhq/xephon-b/issues/36 is solved
		// when the server is overwhelmed, it's pretty likely to have tons of errors ...
		//log.Debugf("%d %s", res.StatusCode, string(b))
		return errors.New(string(body))
	}
	c.proto = res.Proto
	c.bytesSendSuccess += payloadSize
	return nil
}
