package entity

import (
	"testing"
	"time"
)

func TestLocationCreation(t *testing.T) {
	location := Location{
		City:    "New York",
		Country: "US",
		Lat:     40.7128,
		Lon:     -74.0060,
	}

	if location.City != "New York" {
		t.Errorf("Expected City 'New York', got '%s'", location.City)
	}

	if location.Country != "US" {
		t.Errorf("Expected Country 'US', got '%s'", location.Country)
	}

	if location.Lat != 40.7128 {
		t.Errorf("Expected Lat 40.7128, got %f", location.Lat)
	}

	if location.Lon != -74.0060 {
		t.Errorf("Expected Lon -74.0060, got %f", location.Lon)
	}
}

func TestCurrentWeatherCreation(t *testing.T) {
	location := Location{
		City:    "London",
		Country: "UK",
		Lat:     51.5074,
		Lon:     -0.1278,
	}

	condition := WeatherCondition{
		Code:        800,
		Main:        "Clear",
		Description: "clear sky",
		Icon:        "01d",
	}

	weather := CurrentWeather{
		Location:      location,
		Temperature:   15.5,
		FeelsLike:     14.2,
		Humidity:      65,
		Pressure:      1013,
		WindSpeed:     3.5,
		WindDirection: 180,
		Cloudiness:    0,
		Visibility:    10000,
		Condition:     condition,
		Timestamp:     time.Now(),
	}

	if weather.Temperature != 15.5 {
		t.Errorf("Expected Temperature 15.5, got %f", weather.Temperature)
	}

	if weather.Humidity != 65 {
		t.Errorf("Expected Humidity 65, got %d", weather.Humidity)
	}

	if weather.Condition.Main != "Clear" {
		t.Errorf("Expected Condition.Main 'Clear', got '%s'", weather.Condition.Main)
	}
}

func TestDailyForecastCreation(t *testing.T) {
	condition := WeatherCondition{
		Code:        500,
		Main:        "Rain",
		Description: "light rain",
		Icon:        "10d",
	}

	forecast := DailyForecast{
		Date:          time.Now().Add(24 * time.Hour),
		TempMin:       10.0,
		TempMax:       18.0,
		TempMorning:   11.0,
		TempDay:       17.0,
		TempEvening:   14.0,
		TempNight:     12.0,
		Humidity:      70,
		Pressure:      1008,
		WindSpeed:     4.0,
		WindDirection: 220,
		Cloudiness:    60,
		Precipitation: 2.5,
		Condition:     condition,
		Pop:           0.8,
	}

	if forecast.TempMin != 10.0 {
		t.Errorf("Expected TempMin 10.0, got %f", forecast.TempMin)
	}

	if forecast.TempMax != 18.0 {
		t.Errorf("Expected TempMax 18.0, got %f", forecast.TempMax)
	}

	if forecast.Pop != 0.8 {
		t.Errorf("Expected Pop 0.8, got %f", forecast.Pop)
	}

	if forecast.Condition.Main != "Rain" {
		t.Errorf("Expected Condition.Main 'Rain', got '%s'", forecast.Condition.Main)
	}
}

func TestWeatherForecastCreation(t *testing.T) {
	location := Location{
		City:    "Paris",
		Country: "FR",
		Lat:     48.8566,
		Lon:     2.3522,
	}

	forecast := WeatherForecast{
		Location: location,
		Current: CurrentWeather{
			Location:    location,
			Temperature: 12.0,
			Humidity:    75,
		},
		Forecasts: []DailyForecast{
			{
				Date:    time.Now().Add(24 * time.Hour),
				TempMin: 8.0,
				TempMax: 15.0,
			},
			{
				Date:    time.Now().Add(48 * time.Hour),
				TempMin: 9.0,
				TempMax: 16.0,
			},
		},
		GeneratedAt: time.Now(),
	}

	if len(forecast.Forecasts) != 2 {
		t.Errorf("Expected 2 forecasts, got %d", len(forecast.Forecasts))
	}

	if forecast.Location.City != "Paris" {
		t.Errorf("Expected City 'Paris', got '%s'", forecast.Location.City)
	}
}

func TestWeatherReportCreation(t *testing.T) {
	location := Location{
		City:    "Tokyo",
		Country: "JP",
		Lat:     35.6762,
		Lon:     139.6503,
	}

	now := time.Now()
	cachedTime := now.Add(-5 * time.Minute)

	report := WeatherReport{
		ID:          "test-report-123",
		Location:    location,
		Current:     nil,
		Forecast:    nil,
		RequestedAt: now,
		CachedAt:    &cachedTime,
	}

	if report.ID != "test-report-123" {
		t.Errorf("Expected ID 'test-report-123', got '%s'", report.ID)
	}

	if report.CachedAt == nil {
		t.Error("Expected CachedAt to be set")
	}

	if report.RequestedAt.After(report.CachedAt.Add(5 * time.Minute)) {
		t.Error("RequestedAt should be after CachedAt")
	}
}
