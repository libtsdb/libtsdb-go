package config

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
