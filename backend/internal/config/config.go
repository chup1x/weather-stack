package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	Server   serverConfig
	Postgres postgresConfig
	News     newsServiceConfig
}

type serverConfig struct {
	Port string `env:"SERVER_PORT" envDefault:"80"`
}

type postgresConfig struct {
	Host     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	Port     string `env:"POSTGRES_PORT" envDefault:"5432"`
	Name     string `env:"POSTGRES_NAME" envDefault:"postgres"`
	User     string `env:"POSTGRES_USER" envDefault:"postgres"`
	Password string `env:"POSTGRES_PASSWORD" envDefault:"postgres"`
	SSLMode  string `env:"POSTGRES_SSLMODE" envDefault:"disable"`
}

type newsServiceConfig struct {
	Host   string `env:"NEW_SERVICE_HOST" envDefault:"newsapi.org/v2/everything"`
	APIKey string `env:"NEW_SERVICE_API_KEY" envDefault:"0fac40f7dcd34967af176019e1c6a526"`
}

func GetConfig() (*Config, error) {
	_ = godotenv.Load()

	config := &Config{}
	if err := env.Parse(config); err != nil {
		return nil, fmt.Errorf("parse .env file: %w", err)
	}

	return config, nil
}
