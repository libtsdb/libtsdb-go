package config

import "time"

// TODO: move them to each database's own folder

type AkumuliClientConfig struct {
	Addr    string        `yaml:"addr"`
	Timeout time.Duration `yaml:"timeout"`
}

func NewAkumuliClientConfig() *AkumuliClientConfig {
	return &AkumuliClientConfig{
		Addr:    "localhost:8282",
		Timeout: 5 * time.Second,
	}
}

type InfluxdbClientConfig struct {
	Addr     string `yaml:"addr"`
	Database string `yaml:"database"`
}

func NewInfluxdbClientConfig() *InfluxdbClientConfig {
	return &InfluxdbClientConfig{
		Addr:     "http://localhost:8086",
		Database: "libtsdb",
	}
}

type KairosdbClientConfig struct {
	Addr       string        `yaml:"addr"`
	Telnet     bool          `yaml:"telnet"`
	TelnetAddr string        `yaml:"telnetAddr"`
	Timeout    time.Duration `yaml:"timeout"`
}

func NewKairosdbClientConfig() *KairosdbClientConfig {
	return &KairosdbClientConfig{
		Addr:       "http://localhost:8080",
		TelnetAddr: "localhost:4242",
		Timeout:    5 * time.Second,
	}
}

type XephonKClientConfig struct {
	Addr     string `yaml:"addr"`
	Prepare  bool   `yaml:"prepare"`
	Columnar bool   `yaml:"columnar"`
}

func NewXephonkClientConfig() *XephonKClientConfig {
	return &XephonKClientConfig{
		Addr:     "localhost:2334",
		Prepare:  false,
		Columnar: false,
	}
}

type GraphiteClientConfig struct {
	Addr    string        `yaml:"addr"`
	Timeout time.Duration `yaml:"timeout"`
}

func NewGraphiteClientConfig() *GraphiteClientConfig {
	return &GraphiteClientConfig{
		Addr:    "localhost:2003",
		Timeout: 5 * time.Second,
	}
}
