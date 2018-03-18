package config

import "time"

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
