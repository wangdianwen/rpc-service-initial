package mapper

import (
	"time"

	"rpc-service/internal/application/dto"
	"rpc-service/internal/domain/entity"
)

func ToLocationResponse(location entity.Location) dto.LocationResponse {
	return dto.LocationResponse{
		City:    location.City,
		Country: location.Country,
		Lat:     location.Lat,
		Lon:     location.Lon,
	}
}

func ToConditionResponse(condition entity.WeatherCondition) dto.WeatherConditionResponse {
	return dto.WeatherConditionResponse{
		Code:        condition.Code,
		Main:        condition.Main,
		Description: condition.Description,
		Icon:        condition.Icon,
	}
}

func ToCurrentWeatherResponse(weather *entity.CurrentWeather) dto.CurrentWeatherResponse {
	return dto.CurrentWeatherResponse{
		Location:      ToLocationResponse(weather.Location),
		Temperature:   weather.Temperature,
		FeelsLike:     weather.FeelsLike,
		Humidity:      weather.Humidity,
		Pressure:      weather.Pressure,
		WindSpeed:     weather.WindSpeed,
		WindDirection: weather.WindDirection,
		Cloudiness:    weather.Cloudiness,
		Visibility:    weather.Visibility,
		Condition:     ToConditionResponse(weather.Condition),
		Timestamp:     weather.Timestamp.Format(time.RFC3339),
	}
}

func ToDailyForecastResponse(forecast entity.DailyForecast) dto.DailyForecastResponse {
	return dto.DailyForecastResponse{
		Date:          forecast.Date.Format(time.RFC3339),
		TempMin:       forecast.TempMin,
		TempMax:       forecast.TempMax,
		TempMorning:   forecast.TempMorning,
		TempDay:       forecast.TempDay,
		TempEvening:   forecast.TempEvening,
		TempNight:     forecast.TempNight,
		Humidity:      forecast.Humidity,
		Pressure:      forecast.Pressure,
		WindSpeed:     forecast.WindSpeed,
		WindDirection: forecast.WindDirection,
		Cloudiness:    forecast.Cloudiness,
		Precipitation: forecast.Precipitation,
		Condition:     ToConditionResponse(forecast.Condition),
		Pop:           forecast.Pop,
	}
}

func ToWeatherForecastResponse(forecast *entity.WeatherForecast) dto.WeatherForecastResponse {
	responses := make([]dto.DailyForecastResponse, len(forecast.Forecasts))
	for i, f := range forecast.Forecasts {
		responses[i] = ToDailyForecastResponse(f)
	}

	return dto.WeatherForecastResponse{
		Location:    ToLocationResponse(forecast.Location),
		Current:     ToCurrentWeatherResponse(&forecast.Current),
		Forecasts:   responses,
		GeneratedAt: forecast.GeneratedAt.Format(time.RFC3339),
	}
}

func ToLocationEntity(city string, lat, lon float64) entity.Location {
	return entity.Location{
		City: city,
		Lat:  lat,
		Lon:  lon,
	}
}
