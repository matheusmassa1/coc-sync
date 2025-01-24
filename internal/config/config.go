package config

import (
	"time"
)

type Config struct {
	APIKey      string        `env:"API_KEY" env-required:"true"`
	BaseURL     string        `yaml:"base_url" env:"BASE_URL" env-required:"true"`
	HTTPTimeout time.Duration `yaml:"http_timeout" env:"HTTP_TIMEOUT" env-required:"true"`
	ServerPort  int           `yaml:"server_port" env-default:"3000"`
}
