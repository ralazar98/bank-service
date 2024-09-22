package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Env           string `yaml:"env" env-default:"local"`
	HTTPConfig    `yaml:"http_config" `
	StorageConfig `yaml:"storage_config"`
}

type StorageConfig struct {
	Host     string `yaml:"host" `
	Port     uint16 `yaml:"port" `
	Database string `yaml:"database" `
	Username string `yaml:"username" `
	Password string `yaml:"password" `
}

type HTTPConfig struct {
	Address string `yaml:"address"`
}

func NewConfig() *Config {
	return &Config{
		Env: "",
		HTTPConfig: HTTPConfig{
			Address: "",
		},
		StorageConfig: StorageConfig{
			Host:     "",
			Port:     8080,
			Database: "",
			Username: "",
			Password: "",
		},
	}
}

func InitConfig() *Config {
	config := NewConfig()
	file, err := os.Open("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	yamlFile, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
