package config

import "time"

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
