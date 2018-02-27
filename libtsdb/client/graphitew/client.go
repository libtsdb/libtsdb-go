package graphitew

import (
	"net"
	"time"

	"github.com/dyweb/gommon/errors"

	"github.com/libtsdb/libtsdb-go/libtsdb/common"
	"github.com/libtsdb/libtsdb-go/libtsdb/common/graphite"
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
