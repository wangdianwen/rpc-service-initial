package external

import (
	"context"
	"time"

	"rpc-service/internal/domain"
	"rpc-service/internal/domain/entity"
)

type MockWeatherAPIClient struct {
	ShouldFail     bool
	FailError      error
	CustomResponse *entity.CurrentWeather
}

func NewMockWeatherAPIClient() *MockWeatherAPIClient {
	return &MockWeatherAPIClient{
		ShouldFail: false,
		FailError:  nil,
	}
}

func (c *MockWeatherAPIClient) GetCurrentWeather(ctx context.Context, location entity.Location) (*entity.CurrentWeather, error) {
	if c.ShouldFail {
		if c.FailError != nil {
			return nil, c.FailError
		}
		return nil, &domain.DomainError{
			Code:    "MOCK_ERROR",
			Message: "mock error",
			Err:     domain.ErrWeatherAPIError,
		}
	}

	if c.CustomResponse != nil {
		return c.CustomResponse, nil
	}

	return &entity.CurrentWeather{
		Location: entity.Location{
			City: location.City,
			Lat:  location.Lat,
			Lon:  location.Lon,
		},
		Temperature:   20.5,
		FeelsLike:     19.2,
		Humidity:      60,
		Pressure:      1015,
		WindSpeed:     3.0,
		WindDirection: 180,
		Cloudiness:    20,
		Visibility:    10000,
		Condition: entity.WeatherCondition{
			Code:        800,
			Main:        "Clear",
			Description: "clear sky",
			Icon:        "01d",
		},
		Timestamp: time.Now(),
	}, nil
}

func (c *MockWeatherAPIClient) GetForecast(ctx context.Context, location entity.Location, days int) (*entity.WeatherForecast, error) {
	if c.ShouldFail {
		if c.FailError != nil {
			return nil, c.FailError
		}
		return nil, &domain.DomainError{
			Code:    "MOCK_ERROR",
			Message: "mock error",
			Err:     domain.ErrWeatherAPIError,
		}
	}

	forecasts := make([]entity.DailyForecast, 0, days)
	for i := 0; i < days; i++ {
		forecasts = append(forecasts, entity.DailyForecast{
			Date:          time.Now().Add(time.Duration(i+1) * 24 * time.Hour),
			TempMin:       15.0 + float64(i),
			TempMax:       25.0 + float64(i),
			TempMorning:   16.0 + float64(i),
			TempDay:       24.0 + float64(i),
			TempEvening:   20.0 + float64(i),
			TempNight:     17.0 + float64(i),
			Humidity:      60 + i*5,
			Pressure:      1010 + i,
			WindSpeed:     3.0 + float64(i),
			WindDirection: 180 + i*10,
			Cloudiness:    20 + i*10,
			Precipitation: 0.0,
			Condition: entity.WeatherCondition{
				Code:        800,
				Main:        "Clear",
				Description: "clear sky",
				Icon:        "01d",
			},
			Pop: 0.1,
		})
	}

	return &entity.WeatherForecast{
		Location: entity.Location{
			City: location.City,
			Lat:  location.Lat,
			Lon:  location.Lon,
		},
		Current: entity.CurrentWeather{
			Location: entity.Location{
				City: location.City,
				Lat:  location.Lat,
				Lon:  location.Lon,
			},
			Temperature: 20.5,
			Humidity:    60,
		},
		Forecasts:   forecasts,
		GeneratedAt: time.Now(),
	}, nil
}

type FailingWeatherAPIClient struct{}

func (c *FailingWeatherAPIClient) GetCurrentWeather(ctx context.Context, location entity.Location) (*entity.CurrentWeather, error) {
	return nil, &domain.DomainError{
		Code:    "API_ERROR",
		Message: "simulated API failure",
		Err:     domain.ErrWeatherAPIError,
	}
}

func (c *FailingWeatherAPIClient) GetForecast(ctx context.Context, location entity.Location, days int) (*entity.WeatherForecast, error) {
	return nil, &domain.DomainError{
		Code:    "API_ERROR",
		Message: "simulated API failure",
		Err:     domain.ErrWeatherAPIError,
	}
}
