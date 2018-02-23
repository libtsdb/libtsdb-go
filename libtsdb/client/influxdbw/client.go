package influxdbw

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/dyweb/gommon/requests"
	"github.com/pkg/errors"

	"github.com/libtsdb/libtsdb-go/libtsdb/common/influxdb"
	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"
)

type Config struct {
	Addr     string `yaml:"addr"`
	Database string `yaml:"database"`
}

type Client struct {
	enc     *influxdb.Encoder
	h       *http.Client
	baseReq *http.Request
	baseURL *url.URL
}

func New(cfg Config) (*Client, error) {
	u, err := url.Parse(cfg.Addr)
	if err != nil {
		return nil, errors.Wrap(err, "can't parse server address")
	}
	baseReq, err := http.NewRequest("POST", u.String()+"/write", nil)
	if err != nil {
		return nil, errors.Wrap(err, "can't create base query")
	}
	params := baseReq.URL.Query()
	params.Set("db", cfg.Database)
	baseReq.URL.RawQuery = params.Encode()
	baseReq.Header.Set("User-Agent", "libtsdb")
	c := &Client{
		enc:     influxdb.NewEncoder(),
		h:       requests.NewDefaultClient(),
		baseURL: u,
		baseReq: baseReq,
	}
	return c, nil
}

func (c *Client) WriteIntPoint(p *pb.PointIntTagged) error {
	c.enc.WritePointIntTagged(p)
	return c.send()
}

func (c *Client) WriteDoublePoint(p *pb.PointDoubleTagged) error {
	c.enc.WritePointDoubleTagged(p)
	return c.send()
}

func (c *Client) send() error {
	req := &http.Request{}
	*req = *c.baseReq
	req.Body = c.enc.ReadCloser()
	res, err := c.h.Do(req)
	c.enc.Reset()
	if err != nil {
		return errors.Wrap(err, "error send http request")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.Wrap(err, "can't read response body")
	}
	if res.StatusCode != http.StatusNoContent && res.StatusCode != http.StatusOK {
		return errors.New(string(body))
	}
	return nil
}
