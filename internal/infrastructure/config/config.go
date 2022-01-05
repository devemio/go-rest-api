package config

type Config struct {
}

func New() *Config {
	return &Config{}
}

func (c *Config) GetPort() int {
	return 8080
}
