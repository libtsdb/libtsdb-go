package graphitew

import (
	"github.com/dyweb/gommon/errors"
	"net"
	"time"

	"github.com/libtsdb/libtsdb-go/libtsdb/common"
	"github.com/libtsdb/libtsdb-go/libtsdb/common/graphite"
	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"
)

type Config struct {
	Addr string `yaml:"addr"`
}

// Client is a graphite write client using TCP
// TODO: ref promethus and telegraf
type Client struct {
	enc  common.Encoder
	conn net.Conn
}

func New(cfg Config) (*Client, error) {
	// TODO: allow config timeout
	conn, err := net.DialTimeout("tcp", cfg.Addr, 5*time.Second)
	if err != nil {
		return nil, errors.Wrap(err, "can't dial tcp")
	}
	return &Client{
		enc:  graphite.NewTextEncoder(),
		conn: conn,
	}, nil
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
	_, err := c.conn.Write(c.enc.Bytes())
	c.enc.Reset()
	if err != nil {
		return errors.Wrap(err, "error send http request")
	}
	return nil
}
