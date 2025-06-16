package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type (
	Config struct {
		HTTP HTTP
	}

	HTTP struct {
		Port string `env:"HTTP_PORT,required"`
	}
)

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
