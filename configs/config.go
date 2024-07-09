package configs

import "context"

type Config struct {
	Server ServerConfig
}

func New() *Config {
	var cfg Config
	// For run outside of docker, load .env file
	_ = godotenv.Overload()

	// Parse env to struct
	err := envconfig.Process(context.Background(), &cfg)
	if err != nil {
		return nil
	}

	return &cfg
}

type ServerConfig struct {
	Port string `env:"PORT"`
}
