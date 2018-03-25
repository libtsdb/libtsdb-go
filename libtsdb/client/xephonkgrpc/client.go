package xephonkgrpc

import (
	"github.com/xephonhq/xephon-k/xk/client/grpcclient"
	xkconfig "github.com/xephonhq/xephon-k/xk/config"

	"github.com/libtsdb/libtsdb-go/libtsdb/config"
)

func New(cfg config.XephonKClientConfig) (*grpcclient.Client, error) {
	return grpcclient.New(xkconfig.ClientConfig{
		Addr:     cfg.Addr,
		Prepare:  cfg.Prepare,
		Columnar: cfg.Columnar,
	})
}
