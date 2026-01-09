package service

import (
	"context"
	"errors"

	"rpc-service/internal/application/dto"
	"rpc-service/internal/application/mapper"
	"rpc-service/internal/application/validator"
	"rpc-service/internal/domain"
	"rpc-service/internal/domain/entity"
	"rpc-service/internal/domain/repository"
)

type WeatherService struct {
	repo repository.WeatherRepository
}

func NewWeatherService(repo repository.WeatherRepository) *WeatherService {
	return &WeatherService{repo: repo}
}

func (s *WeatherService) GetCurrentWeather(ctx context.Context, req *dto.GetCurrentWeatherRequest) (dto.CurrentWeatherResponse, error) {
	if err := s.validateGetCurrentWeatherRequest(req); err != nil {
		return dto.CurrentWeatherResponse{}, err
	}

	location := entity.Location{
		City: req.City,
		Lat:  req.Lat,
		Lon:  req.Lon,
	}

	weather, err := s.repo.GetCurrentWeather(ctx, location)
	if err != nil {
		if errors.Is(err, domain.ErrWeatherNotFound) {
			return dto.CurrentWeatherResponse{}, &domain.NotFoundError{Resource: "weather", ID: req.City}
		}
		return dto.CurrentWeatherResponse{}, err
	}

	return mapper.ToCurrentWeatherResponse(weather), nil
}

func (s *WeatherService) GetForecast(ctx context.Context, req *dto.GetForecastRequest) (dto.WeatherForecastResponse, error) {
	if err := s.validateGetForecastRequest(req); err != nil {
		return dto.WeatherForecastResponse{}, err
	}

	days := req.Days
	if days < 1 || days > 7 {
		days = 5
	}

	location := entity.Location{
		City: req.City,
		Lat:  req.Lat,
		Lon:  req.Lon,
	}

	forecast, err := s.repo.GetForecast(ctx, location, days)
	if err != nil {
		if errors.Is(err, domain.ErrWeatherNotFound) {
			return dto.WeatherForecastResponse{}, &domain.NotFoundError{Resource: "weather_forecast", ID: req.City}
		}
		return dto.WeatherForecastResponse{}, err
	}

	return mapper.ToWeatherForecastResponse(forecast), nil
}

func (s *WeatherService) validateGetCurrentWeatherRequest(req *dto.GetCurrentWeatherRequest) error {
	result := validator.ValidateGetCurrentWeatherRequest(req.City, req.Lat, req.Lon)
	if !result.Valid {
		return &domain.ValidationError{
			Field:   "request",
			Message: result.Errors[0],
		}
	}
	return nil
}

func (s *WeatherService) validateGetForecastRequest(req *dto.GetForecastRequest) error {
	result := validator.ValidateGetForecastRequest(req.City, req.Lat, req.Lon, req.Days)
	if !result.Valid {
		return &domain.ValidationError{
			Field:   "request",
			Message: result.Errors[0],
		}
	}
	return nil
}
