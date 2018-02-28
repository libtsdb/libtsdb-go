package influxdbr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/dyweb/gommon/errors"
	"github.com/dyweb/gommon/requests"

	influx "github.com/influxdata/influxdb/client/v2"
)

type Config struct {
	Addr     string `yaml:"addr"`
	Database string `yaml:"database"`
}

type Client struct {
	h       *http.Client
	baseReq *http.Request
	baseURL *url.URL
}

// TODO: query does not have chunked, chunksize, params, precision etc.
type Query struct {
	Command  string
	Database string
}

func New(cfg Config) (*Client, error) {
	u, err := url.Parse(cfg.Addr)
	if err != nil {
		return nil, errors.Wrap(err, "can't parse server address")
	}
	baseReq, err := http.NewRequest("POST", u.String()+"/query", nil)
	if err != nil {
		return nil, errors.Wrap(err, "can't create base query")
	}
	params := baseReq.URL.Query()
	params.Set("db", cfg.Database)
	baseReq.URL.RawQuery = params.Encode()
	baseReq.Header.Set("User-Agent", "libtsdb")
	c := &Client{
		h:       requests.NewDefaultClient(),
		baseURL: u,
		baseReq: baseReq,
	}
	return c, nil
}

// curl -XPOST "http://localhost:8086/query" --data-urlencode "q=CREATE DATABASE mydb"
func (c *Client) CreateDatabase(name string) (*influx.Response, error) {
	q := Query{
		Command:  fmt.Sprintf("CREATE DATABASE %s", name),
		Database: "",
	}
	return c.Query(q)
}

// TODO: params and chunked is not supported
func (c *Client) Query(q Query) (*influx.Response, error) {
	req := &http.Request{}
	*req = *c.baseReq
	params := req.URL.Query()
	params.Set("q", q.Command)
	// db is set via config when create client, but we remove it when create database
	if q.Database == "" {
		params.Del("db")
	}
	req.URL.RawQuery = params.Encode()
	res, err := c.h.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error send http request")
	}
	defer res.Body.Close()

	var response influx.Response
	dec := json.NewDecoder(res.Body)
	dec.UseNumber()
	if err := dec.Decode(&response); err != nil {
		if err.Error() == "EOF" && res.StatusCode != http.StatusOK {
			// ignore it, it's not error in json format
		} else {
			return nil, errors.Wrap(err, "error decode")
		}
	}
	// NOTE: this is from influx's code, a non 200 code with a valid response content is not an error
	if res.StatusCode != http.StatusOK && response.Error() == nil {
		return &response, errors.Errorf("error status code %d", res.StatusCode)
	}
	return &response, nil
}
