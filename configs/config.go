package configs

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	App      AppConfig
	Database DataBaseConfig
	RabbitMQ RabbitMQConfig
}

type AppConfig struct {
	Port       string `mapstructure:"port"`
	AppVersion string `mapstructure:"appVersion"`
}

type DataBaseConfig struct {
	Host           string `mapstructure:"host"`
	Port           int    `mapstructure:"port"`
	DBName         string `mapstructure:"dbname"`
	User           string `mapstructure:"user"`
	Password       string `mapstructure:"password"`
	MaxConnections int    `mapstructure:"maxconnections"`
}

type RabbitMQConfig struct {
	RabbitURL   string `mapstructure:"rabbitURL"`
	NameOfQueue string `mapstructure:"nameOfQueue"`
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Error reading config file, %s", err)
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Printf("unable to decode into config struct, %v", err)
	}
	return &config, nil
}
