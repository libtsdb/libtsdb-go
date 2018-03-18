package config

import "time"

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
