package entity

import (
	"time"
)

type Location struct {
	City    string  `json:"city"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}

type WeatherCondition struct {
	Code        int    `json:"code"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type CurrentWeather struct {
	Location      Location         `json:"location"`
	Temperature   float64          `json:"temperature"`
	FeelsLike     float64          `json:"feels_like"`
	Humidity      int              `json:"humidity"`
	Pressure      int              `json:"pressure"`
	WindSpeed     float64          `json:"wind_speed"`
	WindDirection int              `json:"wind_direction"`
	Cloudiness    int              `json:"cloudiness"`
	Visibility    int              `json:"visibility"`
	Condition     WeatherCondition `json:"condition"`
	Timestamp     time.Time        `json:"timestamp"`
}

type DailyForecast struct {
	Date          time.Time        `json:"date"`
	TempMin       float64          `json:"temp_min"`
	TempMax       float64          `json:"temp_max"`
	TempMorning   float64          `json:"temp_morning"`
	TempDay       float64          `json:"temp_day"`
	TempEvening   float64          `json:"temp_evening"`
	TempNight     float64          `json:"temp_night"`
	Humidity      int              `json:"humidity"`
	Pressure      int              `json:"pressure"`
	WindSpeed     float64          `json:"wind_speed"`
	WindDirection int              `json:"wind_direction"`
	Cloudiness    int              `json:"cloudiness"`
	Precipitation float64          `json:"precipitation"`
	Condition     WeatherCondition `json:"condition"`
	Pop           float64          `json:"pop"`
}

type WeatherForecast struct {
	Location    Location        `json:"location"`
	Current     CurrentWeather  `json:"current"`
	Forecasts   []DailyForecast `json:"forecasts"`
	GeneratedAt time.Time       `json:"generated_at"`
}

type WeatherReport struct {
	ID          string
	Location    Location
	Current     *CurrentWeather
	Forecast    *WeatherForecast
	RequestedAt time.Time
	CachedAt    *time.Time
}
