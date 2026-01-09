package configuration

import (
	"fmt"
	"os"
)

type Config struct {
	Port           int    `mapstructure:"port"`
	Host           string `mapstructure:"host"`
	Env            string `mapstructure:"env"`
	LogLevel       string `mapstructure:"log_level"`
	WeatherAPIKey  string `mapstructure:"weather_api_key"`
	WeatherAPIURL  string `mapstructure:"weather_api_url"`
	WeatherTimeout int    `mapstructure:"weather_timeout"`
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
	weatherAPIKey := ""
	weatherAPIURL := ""
	weatherTimeout := 10

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

	if v := os.Getenv("WEATHER_API_KEY"); v != "" {
		weatherAPIKey = v
	}

	if v := os.Getenv("WEATHER_API_URL"); v != "" {
		weatherAPIURL = v
	}

	if v := os.Getenv("WEATHER_TIMEOUT"); v != "" {
		if _, err := fmt.Sscanf(v, "%d", &weatherTimeout); err != nil {
			return nil, err
		}
	}

	return &Config{
		Port:           port,
		Host:           host,
		Env:            env,
		LogLevel:       logLevel,
		WeatherAPIKey:  weatherAPIKey,
		WeatherAPIURL:  weatherAPIURL,
		WeatherTimeout: weatherTimeout,
	}, nil
}
