package kairosdbw

import (
	"net/http"
	"net/url"

	"github.com/dyweb/gommon/errors"
	"github.com/libtsdb/libtsdb-go/libtsdb/client/genericw"
	"github.com/libtsdb/libtsdb-go/libtsdb/common/kairosdb"
)

type Config struct {
	Addr string `yaml:"addr"`
}

func New(cfg Config) (*genericw.Client, error) {
	u, err := url.Parse(cfg.Addr)
	if err != nil {
		return nil, errors.Wrap(err, "can't parse server address")
	}
	baseReq, err := http.NewRequest("POST", u.String()+"/api/v1/datapoints", nil)
	if err != nil {
		return nil, errors.Wrap(err, "can't create base query")
	}
	c := genericw.New(kairosdb.NewJsonEncoder(), baseReq)
	return c, nil
}
