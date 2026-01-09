package service

import (
	"context"
	"testing"

	"rpc-service/internal/application/dto"
	"rpc-service/internal/domain"
	"rpc-service/internal/infrastructure/external"
)

func TestWeatherService_GetCurrentWeather_Success(t *testing.T) {
	mockClient := external.NewMockWeatherAPIClient()
	repo := external.NewWeatherRepositoryWithCache(mockClient, 15*60*1000000000)
	service := NewWeatherService(repo)

	req := &dto.GetCurrentWeatherRequest{
		City: "London",
		Lat:  0,
		Lon:  0,
	}

	result, err := service.GetCurrentWeather(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if result.Location.City != "London" {
		t.Errorf("Expected City 'London', got '%s'", result.Location.City)
	}

	if result.Temperature == 0 {
		t.Error("Expected non-zero temperature")
	}
}

func TestWeatherService_GetCurrentWeather_ValidationError(t *testing.T) {
	mockClient := external.NewMockWeatherAPIClient()
	repo := external.NewWeatherRepositoryWithCache(mockClient, 15*60*1000000000)
	service := NewWeatherService(repo)

	req := &dto.GetCurrentWeatherRequest{
		City: "",
		Lat:  0,
		Lon:  0,
	}

	_, err := service.GetCurrentWeather(context.Background(), req)
	if err == nil {
		t.Error("Expected validation error for empty city and coordinates")
	}

	_, ok := err.(*domain.ValidationError)
	if !ok {
		t.Errorf("Expected ValidationError, got %T", err)
	}
}

func TestWeatherService_GetForecast_Success(t *testing.T) {
	mockClient := external.NewMockWeatherAPIClient()
	repo := external.NewWeatherRepositoryWithCache(mockClient, 15*60*1000000000)
	service := NewWeatherService(repo)

	req := &dto.GetForecastRequest{
		City: "Paris",
		Lat:  0,
		Lon:  0,
		Days: 5,
	}

	result, err := service.GetForecast(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if result.Location.City != "Paris" {
		t.Errorf("Expected City 'Paris', got '%s'", result.Location.City)
	}

	if len(result.Forecasts) != 5 {
		t.Errorf("Expected 5 forecasts, got %d", len(result.Forecasts))
	}
}

func TestWeatherService_GetForecast_DefaultDays(t *testing.T) {
	mockClient := external.NewMockWeatherAPIClient()
	repo := external.NewWeatherRepositoryWithCache(mockClient, 15*60*1000000000)
	service := NewWeatherService(repo)

	req := &dto.GetForecastRequest{
		City: "Berlin",
		Lat:  0,
		Lon:  0,
		Days: 0,
	}

	result, err := service.GetForecast(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if len(result.Forecasts) != 5 {
		t.Errorf("Expected 5 forecasts (default), got %d", len(result.Forecasts))
	}
}

func TestWeatherService_CacheFunctionality(t *testing.T) {
	mockClient := external.NewMockWeatherAPIClient()
	repo := external.NewWeatherRepositoryWithCache(mockClient, 15*60*1000000000)
	service := NewWeatherService(repo)

	req := &dto.GetCurrentWeatherRequest{
		City: "Tokyo",
		Lat:  0,
		Lon:  0,
	}

	_, err := service.GetCurrentWeather(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	_, err = service.GetCurrentWeather(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected cached result, got error: %v", err)
	}
}

func TestWeatherService_WithFailingClient(t *testing.T) {
	failingClient := &external.FailingWeatherAPIClient{}
	repo := external.NewWeatherRepositoryWithCache(failingClient, 15*60*1000000000)
	service := NewWeatherService(repo)

	req := &dto.GetCurrentWeatherRequest{
		City: "Nowhere",
		Lat:  0,
		Lon:  0,
	}

	_, err := service.GetCurrentWeather(context.Background(), req)
	if err == nil {
		t.Error("Expected error from failing client")
	}
}
