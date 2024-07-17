package config

import "context"

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Port string `env:"PORT" default:"8080"`
}

func New() *Config {
	var cfg Config
	_ = godotenv.Overload()

	// Parse env to struct
	err := envconfig.Process(context.Background(), &cfg)
	if err != nil {
		return nil
	}
	return &cfg
}
