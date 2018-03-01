package genericw

import (
	"io/ioutil"
	"net/http"

	"github.com/dyweb/gommon/errors"
	"github.com/dyweb/gommon/requests"
	"github.com/libtsdb/libtsdb-go/libtsdb/common"
	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"
)

const (
	userAgent = "libtsdb-generic"
)

// Client is a generic HTTP based client for write, it is not go routine safe because encoder
// TODO: allow insecure, because we have https server with self signed certs
type Client struct {
	enc     common.Encoder
	h       *http.Client
	baseReq *http.Request
	agent   string

	// stat collected by client after it started running
	proto              string
	bytesSend          uint64
	bytesSendSuccess   uint64
	intPointWritten    uint64
	doublePointWritten uint64
}

// TODO: pass baseReq etc. normally it should be called by other tsdb client ...
func New(encoder common.Encoder, req *http.Request) *Client {
	return &Client{
		enc:     encoder,
		h:       requests.NewDefaultClient(),
		baseReq: req,
		// TODO: user agent is not used
		agent: userAgent,
	}
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

func (c *Client) send() error {
	payloadSize := uint64(c.enc.Len())
	c.bytesSend += payloadSize

	// TODO: go support http client tracing, we can also use open tracing here ...
	// TODO: real bytes send also include header etc, which we didn't take into account of bytes send
	req := &http.Request{}
	*req = *c.baseReq
	req.Body = c.enc.ReadCloser()
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
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent {
		return errors.New(string(body))
	}

	c.proto = res.Proto
	c.bytesSendSuccess += payloadSize
	return nil
}
