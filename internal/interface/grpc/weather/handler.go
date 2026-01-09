package weather

import (
	"context"

	"rpc-service/internal/application/dto"
	"rpc-service/internal/application/service"
	pb "rpc-service/pb/weather"
)

type WeatherServer struct {
	pb.UnimplementedWeatherServiceServer
	weatherService *service.WeatherService
}

func NewWeatherServer(weatherService *service.WeatherService) *WeatherServer {
	return &WeatherServer{
		weatherService: weatherService,
	}
}

func (s *WeatherServer) GetCurrentWeather(ctx context.Context, req *pb.CurrentWeatherRequest) (*pb.CurrentWeatherResponse, error) {
	dtoReq := &dto.GetCurrentWeatherRequest{
		City: req.City,
		Lat:  req.Lat,
		Lon:  req.Lon,
	}

	weather, err := s.weatherService.GetCurrentWeather(ctx, dtoReq)
	if err != nil {
		return &pb.CurrentWeatherResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &pb.CurrentWeatherResponse{
		Success: true,
		Data: &pb.CurrentWeather{
			Location: &pb.Location{
				City:    weather.Location.City,
				Country: weather.Location.Country,
				Lat:     weather.Location.Lat,
				Lon:     weather.Location.Lon,
			},
			Temperature:   weather.Temperature,
			FeelsLike:     weather.FeelsLike,
			Humidity:      int32(weather.Humidity),
			Pressure:      int32(weather.Pressure),
			WindSpeed:     weather.WindSpeed,
			WindDirection: int32(weather.WindDirection),
			Cloudiness:    int32(weather.Cloudiness),
			Visibility:    int32(weather.Visibility),
			Condition: &pb.WeatherCondition{
				Code:        int32(weather.Condition.Code),
				Main:        weather.Condition.Main,
				Description: weather.Condition.Description,
				Icon:        weather.Condition.Icon,
			},
			Timestamp: weather.Timestamp,
		},
		Error: "",
	}, nil
}

func (s *WeatherServer) GetForecast(ctx context.Context, req *pb.ForecastRequest) (*pb.ForecastResponse, error) {
	dtoReq := &dto.GetForecastRequest{
		City: req.City,
		Lat:  req.Lat,
		Lon:  req.Lon,
		Days: int(req.Days),
	}

	forecast, err := s.weatherService.GetForecast(ctx, dtoReq)
	if err != nil {
		return &pb.ForecastResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	pbForecasts := make([]*pb.DailyForecast, len(forecast.Forecasts))
	for i, f := range forecast.Forecasts {
		pbForecasts[i] = &pb.DailyForecast{
			Date:          f.Date,
			TempMin:       f.TempMin,
			TempMax:       f.TempMax,
			TempMorning:   f.TempMorning,
			TempDay:       f.TempDay,
			TempEvening:   f.TempEvening,
			TempNight:     f.TempNight,
			Humidity:      int32(f.Humidity),
			Pressure:      int32(f.Pressure),
			WindSpeed:     f.WindSpeed,
			WindDirection: int32(f.WindDirection),
			Cloudiness:    int32(f.Cloudiness),
			Precipitation: f.Precipitation,
			Condition: &pb.WeatherCondition{
				Code:        int32(f.Condition.Code),
				Main:        f.Condition.Main,
				Description: f.Condition.Description,
				Icon:        f.Condition.Icon,
			},
			Pop: f.Pop,
		}
	}

	return &pb.ForecastResponse{
		Success: true,
		Data: &pb.Forecast{
			Location: &pb.Location{
				City:    forecast.Location.City,
				Country: forecast.Location.Country,
				Lat:     forecast.Location.Lat,
				Lon:     forecast.Location.Lon,
			},
			Current: &pb.CurrentWeather{
				Location: &pb.Location{
					City:    forecast.Current.Location.City,
					Country: forecast.Current.Location.Country,
					Lat:     forecast.Current.Location.Lat,
					Lon:     forecast.Current.Location.Lon,
				},
				Temperature:   forecast.Current.Temperature,
				FeelsLike:     forecast.Current.FeelsLike,
				Humidity:      int32(forecast.Current.Humidity),
				Pressure:      int32(forecast.Current.Pressure),
				WindSpeed:     forecast.Current.WindSpeed,
				WindDirection: int32(forecast.Current.WindDirection),
				Cloudiness:    int32(forecast.Current.Cloudiness),
				Visibility:    int32(forecast.Current.Visibility),
				Condition: &pb.WeatherCondition{
					Code:        int32(forecast.Current.Condition.Code),
					Main:        forecast.Current.Condition.Main,
					Description: forecast.Current.Condition.Description,
					Icon:        forecast.Current.Condition.Icon,
				},
				Timestamp: forecast.Current.Timestamp,
			},
			Forecasts:   pbForecasts,
			GeneratedAt: forecast.GeneratedAt,
		},
		Error: "",
	}, nil
}
