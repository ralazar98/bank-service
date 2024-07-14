package config

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Port string `default:"8080"`
}

func New() *Config {
	var cfg Config

	return &cfg
}
