package graphitew

import (
	"github.com/libtsdb/libtsdb-go/libtsdb/client/genericw"
	"github.com/libtsdb/libtsdb-go/libtsdb/common/graphite"
	"github.com/libtsdb/libtsdb-go/libtsdb/config"
)

func New(cfg config.GraphiteClientConfig) (*genericw.TcpClient, error) {
	return genericw.NewTcp(graphite.Meta(), graphite.NewTextEncoder(), cfg.Addr, cfg.Timeout)
}
