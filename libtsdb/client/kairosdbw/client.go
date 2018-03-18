package kairosdbw

import (
	"net/http"
	"net/url"

	"github.com/dyweb/gommon/errors"

	"github.com/libtsdb/libtsdb-go/libtsdb/client/genericw"
	"github.com/libtsdb/libtsdb-go/libtsdb/common/kairosdb"
	"github.com/libtsdb/libtsdb-go/libtsdb/config"
)

func NewHttp(cfg config.KairosdbClientConfig) (*genericw.HttpClient, error) {
	u, err := url.Parse(cfg.Addr)
	if err != nil {
		return nil, errors.Wrap(err, "can't parse server address")
	}
	baseReq, err := http.NewRequest("POST", u.String()+"/api/v1/datapoints", nil)
	if err != nil {
		return nil, errors.Wrap(err, "can't create base query")
	}
	c := genericw.NewHttp(kairosdb.Meta(), kairosdb.NewJsonEncoder(), baseReq)
	return c, nil
}

func NewTcp(cfg config.KairosdbClientConfig) (*genericw.TcpClient, error) {
	return genericw.NewTcp(kairosdb.Meta(), kairosdb.NewTelnetEncoder(), cfg.TelnetAddr, cfg.Timeout)
}
