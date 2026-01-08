package configuration

import (
	"fmt"
	"os"
)

type Config struct {
	Port     int    `mapstructure:"port"`
	Host     string `mapstructure:"host"`
	Env      string `mapstructure:"env"`
	LogLevel string `mapstructure:"log_level"`
}

type ConfigProvider interface {
	GetConfig() (*Config, error)
}

type EnvConfigProvider struct{}

func NewEnvConfigProvider() *EnvConfigProvider {
	return &EnvConfigProvider{}
}

func (p *EnvConfigProvider) GetConfig() (*Config, error) {
	port := 1234
	host := "0.0.0.0"
	env := "development"
	logLevel := "info"

	if v := os.Getenv("PORT"); v != "" {
		if _, err := fmt.Sscanf(v, "%d", &port); err != nil {
			return nil, err
		}
	}

	if v := os.Getenv("HOST"); v != "" {
		host = v
	}

	if v := os.Getenv("ENV"); v != "" {
		env = v
	}

	if v := os.Getenv("LOG_LEVEL"); v != "" {
		logLevel = v
	}

	return &Config{
		Port:     port,
		Host:     host,
		Env:      env,
		LogLevel: logLevel,
	}, nil
}
