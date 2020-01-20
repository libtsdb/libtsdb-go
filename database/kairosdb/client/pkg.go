package client

import (
	"net/http"
	"net/url"

	"github.com/dyweb/gommon/errors"
	"github.com/libtsdb/libtsdb-go/database/kairosdb/config"
	"github.com/libtsdb/libtsdb-go/protocol"
)

func NewKairosDBHTTPClient(cfg config.KairosdbClientConfig) (*protocol.HTTPClient, error) {
	u, err := url.Parse(cfg.Addr)
	if err != nil {
		return nil, errors.Wrap(err, "can't parse server address")
	}
	baseReq, err := http.NewRequest("POST", u.String()+"/api/v1/datapoints", nil)
	if err != nil {
		return nil, errors.Wrap(err, "can't create base query")
	}
	c := protocol.NewHTTPClient(NewKairosDBJSONEncoder(), baseReq)
	return c, nil
}

func NewKairosDBTCPClient(cfg config.KairosdbClientConfig) (*protocol.TCPClient, error) {
	return protocol.NewTCPClient(NewKaiorsDBTelnetEncoder(), cfg.TelnetAddr, cfg.Timeout)
}
