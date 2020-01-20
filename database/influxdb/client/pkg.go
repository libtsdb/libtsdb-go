package client

import (
	"net/http"
	"net/url"

	"github.com/dyweb/gommon/errors"
	"github.com/libtsdb/libtsdb-go/database"
	"github.com/libtsdb/libtsdb-go/database/influxdb"
	"github.com/libtsdb/libtsdb-go/database/influxdb/config"
	"github.com/libtsdb/libtsdb-go/protocol"
)

type InfluxDBClient struct {
	*protocol.HTTPClient
	meta database.Meta
}

func (c *InfluxDBClient) Meta() database.Meta {
	return c.meta
}

func NewInfluxDBClient(cfg config.InfluxdbClientConfig) (*InfluxDBClient, error) {
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
	return &InfluxDBClient{HTTPClient: c, meta: influxdb.Meta()}, nil
}
