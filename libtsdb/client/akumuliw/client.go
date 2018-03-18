package akumuliw

import (
	"github.com/libtsdb/libtsdb-go/libtsdb/client/genericw"
	"github.com/libtsdb/libtsdb-go/libtsdb/common/akumuli"
	"github.com/libtsdb/libtsdb-go/libtsdb/config"
)

func New(cfg config.AkumuliClientConfig) (*genericw.TcpClient, error) {
	return genericw.NewTcp(akumuli.Meta(), akumuli.NewEncoder(), cfg.Addr, cfg.Timeout)
}
