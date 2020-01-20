package client

import (
	"net/http"
	"net/url"

	"github.com/dyweb/gommon/errors"
	"github.com/libtsdb/libtsdb-go/database/influxdb/config"
	"github.com/libtsdb/libtsdb-go/protocol"
)

func New(cfg config.InfluxdbClientConfig) (*protocol.HTTPClient, error) {
	u, err := url.Parse(cfg.Addr)
	if err != nil {
		return nil, errors.Wrap(err, "can't parse server address")
	}
	baseReq, err := http.NewRequest("POST", u.String()+"/write", nil)
	if err != nil {
		return nil, errors.Wrap(err, "can't create base query")
	}
	params := baseReq.URL.Query()
	params.Set("db", cfg.Database)
	baseReq.URL.RawQuery = params.Encode()
	baseReq.Header.Set("User-Agent", "libtsdb")
	c := protocol.NewHTTPClient(NewInfluxDBEncoder(), baseReq)
	return c, nil
}
