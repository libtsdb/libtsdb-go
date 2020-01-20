package libtsdb

import (
	"github.com/dyweb/gommon/errors"
	"github.com/libtsdb/libtsdb-go/database"
	"github.com/libtsdb/libtsdb-go/database/akumuli/client"
	akumuliconfig "github.com/libtsdb/libtsdb-go/database/akumuli/config"
	influxdbclient "github.com/libtsdb/libtsdb-go/database/influxdb/client"
	influxdbconfig "github.com/libtsdb/libtsdb-go/database/influxdb/config"
	kairosdbclient "github.com/libtsdb/libtsdb-go/database/kairosdb/client"
	kairosdbconfig "github.com/libtsdb/libtsdb-go/database/kairosdb/config"
)

type DatabaseConfig struct {
	Name     string                               `yaml:"name"`
	Type     string                               `yaml:"type"`
	Akumuli  *akumuliconfig.AkumuliClientConfig   `yaml:"akumuli"`
	Influxdb *influxdbconfig.InfluxdbClientConfig `yaml:"influxdb"`
	Kairosdb *kairosdbconfig.KairosdbClientConfig `yaml:"kairosdb"`
}

func CreateClient(cfg DatabaseConfig) (database.TracedWriteClient, error) {
	switch cfg.Type {
	case "akumuli":
		if cfg.Akumuli == nil {
			return nil, errors.New("akumuli is selected but no config")
		}
		return client.NewAkumuliClient(*cfg.Akumuli)
	case "influxdb":
		if cfg.Influxdb == nil {
			return nil, errors.New("influxdb is selected but no config")
		}
		return influxdbclient.NewInfluxDBClient(*cfg.Influxdb)
	case "kairosdb":
		if cfg.Kairosdb == nil {
			return nil, errors.New("kairosdb is selected but no config")
		}
		if cfg.Kairosdb.Telnet {
			return kairosdbclient.NewKairosDBTCPClient(*cfg.Kairosdb)
		}
		return kairosdbclient.NewKairosDBHTTPClient(*cfg.Kairosdb)
	default:
		return nil, errors.Errorf("unknown database %s", cfg.Type)
	}
}
