package graphitew

import (
	"net"

	"github.com/dyweb/gommon/errors"

	"github.com/libtsdb/libtsdb-go/libtsdb"
	"github.com/libtsdb/libtsdb-go/libtsdb/common"
	"github.com/libtsdb/libtsdb-go/libtsdb/common/graphite"
	"github.com/libtsdb/libtsdb-go/libtsdb/config"
	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"
)

var _ libtsdb.WriteClient = (*Client)(nil)

// Client is a graphite write client using TCP
type Client struct {
	cfg  config.GraphiteClientConfig
	enc  common.Encoder
	conn net.Conn
}

func New(cfg config.GraphiteClientConfig) (*Client, error) {
	conn, err := net.DialTimeout("tcp", cfg.Addr, cfg.Timeout)
	if err != nil {
		return nil, errors.Wrap(err, "can't dial tcp")
	}
	return &Client{
		cfg:  cfg,
		enc:  graphite.NewTextEncoder(),
		conn: conn,
	}, nil
}

func (c *Client) Meta() libtsdb.Meta {
	return graphite.Meta()
}

func (c *Client) WriteIntPoint(p *pb.PointIntTagged) {
	c.enc.WritePointIntTagged(p)
}

func (c *Client) WriteDoublePoint(p *pb.PointDoubleTagged) {
	c.enc.WritePointDoubleTagged(p)
}

func (c *Client) WriteSeriesIntTagged(p *pb.SeriesIntTagged) {
	c.enc.WriteSeriesIntTagged(p)
}

func (c *Client) WriteSeriesDoubleTagged(p *pb.SeriesDoubleTagged) {
	c.enc.WriteSeriesDoubleTagged(p)
}

func (c *Client) Flush() error {
	return c.send()
}

func (c *Client) Close() error {
	if err := c.conn.Close(); err != nil {
		return errors.Wrap(err, "error closing tcp connection")
	}
	return nil
}

func (c *Client) send() error {
	_, err := c.conn.Write(c.enc.Bytes())
	c.enc.Reset()
	if err != nil {
		return errors.Wrap(err, "error send http request")
	}
	return nil
}
