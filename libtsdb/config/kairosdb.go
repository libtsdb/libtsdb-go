package config

type KairosdbClientConfig struct {
	Addr string `yaml:"addr"`
}

func NewKaiorsdbClientConfig() *KairosdbClientConfig {
	return &KairosdbClientConfig{
		Addr: "http://localhost:8080",
	}
}
