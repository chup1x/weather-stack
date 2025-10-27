package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	Server serverConfig
}

type serverConfig struct {
	RemoteHost string `env:"REMOTE_HOST" envDefault:"http://backend:80/api/v1"`
	Port       string `env:"SERVER_PORT" envDefault:"8080"`
}

func GetConfig() (*Config, error) {
	_ = godotenv.Load()

	config := &Config{}
	if err := env.Parse(config); err != nil {
		return nil, fmt.Errorf("parse .env file: %w", err)
	}

	return config, nil
}
