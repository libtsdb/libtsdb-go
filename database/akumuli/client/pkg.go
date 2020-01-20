package client

import (
	"github.com/libtsdb/libtsdb-go/database"
	"github.com/libtsdb/libtsdb-go/database/akumuli"
	"github.com/libtsdb/libtsdb-go/database/akumuli/config"
	"github.com/libtsdb/libtsdb-go/protocol"
)

type AkumuliClient struct {
	*protocol.TCPClient
	meta database.Meta
}

func (c *AkumuliClient) Meta() database.Meta {
	return c.meta
}

func NewAkumuliClient(cfg config.AkumuliClientConfig) (*AkumuliClient, error) {
	p, err := protocol.NewTCPClient(NewAkumuliEncoder(), cfg.Addr, cfg.Timeout)
	if err != nil {
		return nil, err
	}
	return &AkumuliClient{TCPClient: p, meta: akumuli.Meta()}, nil
}
