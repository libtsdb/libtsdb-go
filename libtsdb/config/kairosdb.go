package config

type KairosdbClientConfig struct {
	Addr string `yaml:"addr"`
}

func NewKairosdbClientConfig() *KairosdbClientConfig {
	return &KairosdbClientConfig{
		Addr: "http://localhost:8080",
	}
}
