package repository

import (
	"context"

	"rpc-service/internal/domain/entity"
)

type WeatherRepository interface {
	GetCurrentWeather(ctx context.Context, location entity.Location) (*entity.CurrentWeather, error)
	GetForecast(ctx context.Context, location entity.Location, days int) (*entity.WeatherForecast, error)
	GetCachedWeather(ctx context.Context, location entity.Location) (*entity.WeatherReport, error)
	CacheWeather(ctx context.Context, report *entity.WeatherReport) error
}
