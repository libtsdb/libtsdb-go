package client

import (
	"net/http"
	"net/url"

	"github.com/dyweb/gommon/errors"
	"github.com/libtsdb/libtsdb-go/database"
	"github.com/libtsdb/libtsdb-go/database/kairosdb"
	"github.com/libtsdb/libtsdb-go/database/kairosdb/config"
	"github.com/libtsdb/libtsdb-go/protocol"
)

type KairosDBHTTPClient struct {
	*protocol.HTTPClient
	meta database.Meta
}

func (c *KairosDBHTTPClient) Meta() database.Meta {
	return c.meta
}

type KairosDBTCPCLient struct {
	*protocol.TCPClient
	meta database.Meta
}

func (c *KairosDBTCPCLient) Meta() database.Meta {
	return c.meta
}

func NewKairosDBHTTPClient(cfg config.KairosdbClientConfig) (*KairosDBHTTPClient, error) {
	u, err := url.Parse(cfg.Addr)
	if err != nil {
		return nil, errors.Wrap(err, "can't parse server address")
	}
	baseReq, err := http.NewRequest("POST", u.String()+"/api/v1/datapoints", nil)
	if err != nil {
		return nil, errors.Wrap(err, "can't create base query")
	}
	c := protocol.NewHTTPClient(NewKairosDBJSONEncoder(), baseReq)
	return &KairosDBHTTPClient{HTTPClient: c, meta: kairosdb.Meta()}, nil
}

func NewKairosDBTCPClient(cfg config.KairosdbClientConfig) (*KairosDBTCPCLient, error) {
	p, err := protocol.NewTCPClient(NewKaiorsDBTelnetEncoder(), cfg.TelnetAddr, cfg.Timeout)
	if err != nil {
		return nil, err
	}
	return &KairosDBTCPCLient{TCPClient: p, meta: kairosdb.Meta()}, nil
}
