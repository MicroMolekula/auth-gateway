package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Server  `yaml:"server"`
	Domains map[string]string `yaml:"domains"`
	CORS    string            `yaml:"cors"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

var Cfg *Config

func NewConfig() (*Config, error) {
	cfg := &Config{}
	yamlPath := os.Getenv("CONFIG_PATH")
	if yamlPath == "" {
		return nil, errors.New("config file environment variable not set")
	}
	yamlFile, err := os.ReadFile(yamlPath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error read config file: %s", err))
	}
	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error parse config file: %s", err))
	}
	return cfg, nil
}

func InitConfig() {
	cfg, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	Cfg = cfg
}
