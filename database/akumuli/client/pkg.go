package client

import (
	"github.com/libtsdb/libtsdb-go/database/akumuli/config"
	"github.com/libtsdb/libtsdb-go/protocol"
)

func NewAkumuliClient(cfg config.AkumuliClientConfig) (*protocol.TCPClient, error) {
	return protocol.NewTCPClient(NewAkumuliEncoder(), cfg.Addr, cfg.Timeout)
}
