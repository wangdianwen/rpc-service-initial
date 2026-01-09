package external

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"rpc-service/internal/domain"
	"rpc-service/internal/domain/entity"
)

const (
	OpenWeatherMapBaseURL     = "https://api.openweathermap.org/data/2.5"
	OpenWeatherMapForecastURL = "https://api.openweathermap.org/data/2.5/forecast"
	DefaultTimeout            = 10 * time.Second
	CacheTTL                  = 15 * time.Minute
)

type WeatherAPIConfig struct {
	APIKey  string
	BaseURL string
	Timeout time.Duration
}

type WeatherAPIResponse struct {
	Coord      Coordinates `json:"coord"`
	Weather    []Weather   `json:"weather"`
	Main       Main        `json:"main"`
	Visibility int         `json:"visibility"`
	Wind       Wind        `json:"wind"`
	Clouds     Clouds      `json:"clouds"`
	DT         int64       `json:"dt"`
	Sys        Sys         `json:"sys"`
	Name       string      `json:"name"`
	Cod        int         `json:"cod"`
	Message    string      `json:"message,omitempty"`
}

type ForecastAPIResponse struct {
	Code    string         `json:"code"`
	Message float64        `json:"message"`
	City    City           `json:"city"`
	List    []ForecastItem `json:"list"`
}

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	SeaLevel  int     `json:"sea_level,omitempty"`
	GrndLevel int     `json:"grnd_level,omitempty"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust,omitempty"`
}

type Clouds struct {
	All int `json:"all"`
}

type Sys struct {
	Country string `json:"country,omitempty"`
	Sunrise int64  `json:"sunrise,omitempty"`
	Sunset  int64  `json:"sunset,omitempty"`
}

type City struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Country    string `json:"country"`
	Population int    `json:"population"`
	Timezone   int    `json:"timezone"`
	Sunrise    int64  `json:"sunrise"`
	Sunset     int64  `json:"sunset"`
}

type ForecastItem struct {
	DT         int64       `json:"dt"`
	Main       Main        `json:"main"`
	Weather    []Weather   `json:"weather"`
	Clouds     Clouds      `json:"clouds"`
	Wind       Wind        `json:"wind"`
	Visibility int         `json:"visibility"`
	Pop        float64     `json:"pop"`
	Rain       Rain        `json:"rain,omitempty"`
	Sys        ForecastSys `json:"sys"`
	DTText     string      `json:"dt_txt"`
}

type Rain struct {
	ThreeH float64 `json:"3h,omitempty"`
}

type ForecastSys struct {
	Pod string `json:"pod"`
}

type WeatherAPIClient interface {
	GetCurrentWeather(ctx context.Context, location entity.Location) (*entity.CurrentWeather, error)
	GetForecast(ctx context.Context, location entity.Location, days int) (*entity.WeatherForecast, error)
}

type weatherAPIClient struct {
	config WeatherAPIConfig
	client *http.Client
}

func NewWeatherAPIClient(config WeatherAPIConfig) WeatherAPIClient {
	if config.Timeout == 0 {
		config.Timeout = DefaultTimeout
	}
	if config.BaseURL == "" {
		config.BaseURL = OpenWeatherMapBaseURL
	}
	return &weatherAPIClient{
		config: config,
		client: &http.Client{
			Timeout: config.Timeout,
		},
	}
}

func (c *weatherAPIClient) GetCurrentWeather(ctx context.Context, location entity.Location) (*entity.CurrentWeather, error) {
	var endpoint string
	if location.City != "" {
		endpoint = fmt.Sprintf("%s/weather?q=%s", c.config.BaseURL, url.QueryEscape(location.City))
	} else {
		endpoint = fmt.Sprintf("%s/weather?lat=%f&lon=%f", c.config.BaseURL, location.Lat, location.Lon)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, &domain.DomainError{
			Code:    "API_REQUEST_FAILED",
			Message: "failed to create request",
			Err:     err,
		}
	}

	req.Header.Set("Accept", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, &domain.DomainError{
			Code:    "API_REQUEST_FAILED",
			Message: "failed to execute request",
			Err:     err,
		}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, &domain.DomainError{
			Code:    "API_RESPONSE_READ_FAILED",
			Message: "failed to read response body",
			Err:     err,
		}
	}

	var weatherResp WeatherAPIResponse
	if err := json.Unmarshal(body, &weatherResp); err != nil {
		return nil, &domain.DomainError{
			Code:    "API_RESPONSE_PARSE_FAILED",
			Message: "failed to parse response",
			Err:     err,
		}
	}

	if resp.StatusCode != http.StatusOK {
		return nil, c.handleAPIError(weatherResp, resp.StatusCode)
	}

	return c.mapToCurrentWeather(weatherResp, location)
}

func (c *weatherAPIClient) GetForecast(ctx context.Context, location entity.Location, days int) (*entity.WeatherForecast, error) {
	if days < 1 || days > 7 {
		days = 5
	}

	var endpoint string
	if location.City != "" {
		endpoint = fmt.Sprintf("%s?q=%s&cnt=%d", OpenWeatherMapForecastURL, url.QueryEscape(location.City), days*8)
	} else {
		endpoint = fmt.Sprintf("%s?lat=%f&lon=%f&cnt=%d", OpenWeatherMapForecastURL, location.Lat, location.Lon, days*8)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, &domain.DomainError{
			Code:    "API_REQUEST_FAILED",
			Message: "failed to create request",
			Err:     err,
		}
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, &domain.DomainError{
			Code:    "API_REQUEST_FAILED",
			Message: "failed to execute request",
			Err:     err,
		}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, &domain.DomainError{
			Code:    "API_RESPONSE_READ_FAILED",
			Message: "failed to read response body",
			Err:     err,
		}
	}

	var forecastResp ForecastAPIResponse
	if err := json.Unmarshal(body, &forecastResp); err != nil {
		return nil, &domain.DomainError{
			Code:    "API_RESPONSE_PARSE_FAILED",
			Message: "failed to parse response",
			Err:     err,
		}
	}

	if resp.StatusCode != http.StatusOK {
		return nil, &domain.DomainError{
			Code:    "API_ERROR",
			Message: fmt.Sprintf("weather API returned status %d: %s", resp.StatusCode, string(body)),
			Err:     domain.ErrWeatherAPIError,
		}
	}

	return c.mapToForecast(forecastResp, location)
}

func (c *weatherAPIClient) handleAPIError(resp WeatherAPIResponse, statusCode int) error {
	if statusCode == http.StatusUnauthorized {
		return &domain.DomainError{
			Code:    "API_UNAUTHORIZED",
			Message: "invalid API key",
			Err:     domain.ErrWeatherAPIError,
		}
	}
	if statusCode == http.StatusTooManyRequests {
		return &domain.DomainError{
			Code:    "API_RATE_LIMITED",
			Message: "rate limit exceeded",
			Err:     domain.ErrRateLimited,
		}
	}
	if statusCode == http.StatusNotFound {
		return &domain.DomainError{
			Code:    "LOCATION_NOT_FOUND",
			Message: fmt.Sprintf("location not found: %s", resp.Message),
			Err:     domain.ErrWeatherNotFound,
		}
	}
	return &domain.DomainError{
		Code:    "API_ERROR",
		Message: fmt.Sprintf("weather API returned status %d", statusCode),
		Err:     domain.ErrWeatherAPIError,
	}
}

func (c *weatherAPIClient) mapToCurrentWeather(resp WeatherAPIResponse, location entity.Location) (*entity.CurrentWeather, error) {
	condition := entity.WeatherCondition{}
	if len(resp.Weather) > 0 {
		condition = entity.WeatherCondition{
			Code:        resp.Weather[0].ID,
			Main:        resp.Weather[0].Main,
			Description: resp.Weather[0].Description,
			Icon:        resp.Weather[0].Icon,
		}
	}

	if location.City == "" {
		location.City = resp.Name
	}

	return &entity.CurrentWeather{
		Location:      location,
		Temperature:   resp.Main.Temp - 273.15,
		FeelsLike:     resp.Main.FeelsLike - 273.15,
		Humidity:      resp.Main.Humidity,
		Pressure:      resp.Main.Pressure,
		WindSpeed:     resp.Wind.Speed,
		WindDirection: resp.Wind.Deg,
		Cloudiness:    resp.Clouds.All,
		Visibility:    resp.Visibility,
		Condition:     condition,
		Timestamp:     time.Unix(resp.DT, 0),
	}, nil
}

func (c *weatherAPIClient) mapToForecast(resp ForecastAPIResponse, location entity.Location) (*entity.WeatherForecast, error) {
	forecasts := make([]entity.DailyForecast, 0, len(resp.List))
	currentTemp := 0.0
	currentFound := false

	for _, item := range resp.List {
		condition := entity.WeatherCondition{}
		if len(item.Weather) > 0 {
			condition = entity.WeatherCondition{
				Code:        item.Weather[0].ID,
				Main:        item.Weather[0].Main,
				Description: item.Weather[0].Description,
				Icon:        item.Weather[0].Icon,
			}
		}

		forecast := entity.DailyForecast{
			Date:          time.Unix(item.DT, 0),
			TempMin:       item.Main.TempMin - 273.15,
			TempMax:       item.Main.TempMax - 273.15,
			TempMorning:   item.Main.Temp - 273.15,
			TempDay:       item.Main.Temp - 273.15,
			TempEvening:   item.Main.Temp - 273.15,
			TempNight:     item.Main.Temp - 273.15,
			Humidity:      item.Main.Humidity,
			Pressure:      item.Main.Pressure,
			WindSpeed:     item.Wind.Speed,
			WindDirection: item.Wind.Deg,
			Cloudiness:    item.Clouds.All,
			Precipitation: item.Rain.ThreeH,
			Condition:     condition,
			Pop:           item.Pop,
		}
		forecasts = append(forecasts, forecast)

		hour := time.Unix(item.DT, 0).Hour()
		if !currentFound && (hour >= 12 && hour <= 14) {
			currentTemp = item.Main.Temp - 273.15
			currentFound = true
		}
	}

	if location.City == "" {
		location.City = resp.City.Name
	}
	location.Country = resp.City.Country

	currentWeather := entity.CurrentWeather{
		Location:    location,
		Temperature: currentTemp,
		Humidity:    0,
		Pressure:    0,
		Timestamp:   time.Now(),
	}

	return &entity.WeatherForecast{
		Location:    location,
		Current:     currentWeather,
		Forecasts:   forecasts,
		GeneratedAt: time.Now(),
	}, nil
}

type CachedWeather struct {
	Report     *entity.WeatherReport
	CachedAt   time.Time
	Expiration time.Time
}

type WeatherRepositoryWithCache struct {
	client   WeatherAPIClient
	cache    map[string]CachedWeather
	cacheTTL time.Duration
}

func NewWeatherRepositoryWithCache(client WeatherAPIClient, ttl time.Duration) *WeatherRepositoryWithCache {
	if ttl == 0 {
		ttl = CacheTTL
	}
	return &WeatherRepositoryWithCache{
		client:   client,
		cache:    make(map[string]CachedWeather),
		cacheTTL: ttl,
	}
}

func (r *WeatherRepositoryWithCache) GetCurrentWeather(ctx context.Context, location entity.Location) (*entity.CurrentWeather, error) {
	cacheKey := r.cacheKey("current", location)
	if cached, ok := r.getFromCache(cacheKey); ok && cached.Report.Current != nil {
		return cached.Report.Current, nil
	}

	weather, err := r.client.GetCurrentWeather(ctx, location)
	if err != nil {
		return nil, err
	}

	report := &entity.WeatherReport{
		Location: location,
		Current:  weather,
	}
	r.cacheReport(cacheKey, report)

	return weather, nil
}

func (r *WeatherRepositoryWithCache) GetForecast(ctx context.Context, location entity.Location, days int) (*entity.WeatherForecast, error) {
	cacheKey := r.cacheKey(fmt.Sprintf("forecast_%d", days), location)
	if cached, ok := r.getFromCache(cacheKey); ok && cached.Report.Forecast != nil {
		return cached.Report.Forecast, nil
	}

	forecast, err := r.client.GetForecast(ctx, location, days)
	if err != nil {
		return nil, err
	}

	report := &entity.WeatherReport{
		Location: location,
		Forecast: forecast,
	}
	r.cacheReport(cacheKey, report)

	return forecast, nil
}

func (r *WeatherRepositoryWithCache) GetCachedWeather(ctx context.Context, location entity.Location) (*entity.WeatherReport, error) {
	cacheKey := r.cacheKey("current", location)
	if cached, ok := r.getFromCache(cacheKey); ok {
		return cached.Report, nil
	}
	return nil, domain.ErrWeatherNotFound
}

func (r *WeatherRepositoryWithCache) CacheWeather(ctx context.Context, report *entity.WeatherReport) error {
	cacheKey := r.cacheKey("current", report.Location)
	r.cacheReport(cacheKey, report)
	return nil
}

func (r *WeatherRepositoryWithCache) cacheKey(prefix string, location entity.Location) string {
	if location.City != "" {
		return fmt.Sprintf("%s:%s", prefix, strings.ToLower(location.City))
	}
	return fmt.Sprintf("%s:%.4f_%.4f", prefix, location.Lat, location.Lon)
}

func (r *WeatherRepositoryWithCache) getFromCache(key string) (CachedWeather, bool) {
	cached, ok := r.cache[key]
	if !ok {
		return CachedWeather{}, false
	}
	if time.Now().After(cached.Expiration) {
		delete(r.cache, key)
		return CachedWeather{}, false
	}
	return cached, true
}

func (r *WeatherRepositoryWithCache) cacheReport(key string, report *entity.WeatherReport) {
	now := time.Now()
	r.cache[key] = CachedWeather{
		Report:     report,
		CachedAt:   now,
		Expiration: now.Add(r.cacheTTL),
	}
}
