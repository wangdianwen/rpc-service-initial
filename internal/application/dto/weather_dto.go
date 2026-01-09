package dto

type GetCurrentWeatherRequest struct {
	City string  `json:"city" validate:"required"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
}

type GetForecastRequest struct {
	City string  `json:"city" validate:"required"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
	Days int     `json:"days"`
}

type LocationResponse struct {
	City    string  `json:"city"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}

type WeatherConditionResponse struct {
	Code        int    `json:"code"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type CurrentWeatherResponse struct {
	Location      LocationResponse         `json:"location"`
	Temperature   float64                  `json:"temperature"`
	FeelsLike     float64                  `json:"feels_like"`
	Humidity      int                      `json:"humidity"`
	Pressure      int                      `json:"pressure"`
	WindSpeed     float64                  `json:"wind_speed"`
	WindDirection int                      `json:"wind_direction"`
	Cloudiness    int                      `json:"cloudiness"`
	Visibility    int                      `json:"visibility"`
	Condition     WeatherConditionResponse `json:"condition"`
	Timestamp     string                   `json:"timestamp"`
}

type DailyForecastResponse struct {
	Date          string                   `json:"date"`
	TempMin       float64                  `json:"temp_min"`
	TempMax       float64                  `json:"temp_max"`
	TempMorning   float64                  `json:"temp_morning"`
	TempDay       float64                  `json:"temp_day"`
	TempEvening   float64                  `json:"temp_evening"`
	TempNight     float64                  `json:"temp_night"`
	Humidity      int                      `json:"humidity"`
	Pressure      int                      `json:"pressure"`
	WindSpeed     float64                  `json:"wind_speed"`
	WindDirection int                      `json:"wind_direction"`
	Cloudiness    int                      `json:"cloudiness"`
	Precipitation float64                  `json:"precipitation"`
	Condition     WeatherConditionResponse `json:"condition"`
	Pop           float64                  `json:"pop"`
}

type WeatherForecastResponse struct {
	Location    LocationResponse        `json:"location"`
	Current     CurrentWeatherResponse  `json:"current"`
	Forecasts   []DailyForecastResponse `json:"forecasts"`
	GeneratedAt string                  `json:"generated_at"`
}

type WeatherErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
