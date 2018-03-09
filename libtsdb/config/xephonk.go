package config

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
