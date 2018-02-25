package graphitew

type Config struct {
	Addr string `yaml:"addr"`
}

// Client is a graphite write client using TCP
// TODO: ref promethus and telegraf
type Client struct {
}
