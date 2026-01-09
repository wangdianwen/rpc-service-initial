## ADDED Requirements

### Requirement: Current Weather Query
The weather gRPC service SHALL provide a GetCurrentWeather RPC method that returns current weather conditions for a specified location.

#### Scenario: Get current weather by city name
- **WHEN** client calls GetCurrentWeather with valid city name
- **THEN** return current temperature, humidity, wind speed, and conditions

#### Scenario: Get current weather by coordinates
- **WHEN** client calls GetCurrentWeather with latitude and longitude
- **THEN** return current weather data for the exact location

#### Scenario: Invalid location provided
- **WHEN** client calls GetCurrentWeather with unknown or invalid location
- **THEN** return NOT_FOUND error with appropriate message

### Requirement: Weather Forecast Query
The weather gRPC service SHALL provide a GetForecast RPC method that returns multi-day weather forecasts.

#### Scenario: Get 5-day forecast
- **WHEN** client calls GetForecast with valid location
- **THEN** return daily forecast for next 5 days with high/low temps and conditions

#### Scenario: Forecast with custom days
- **WHEN** client calls GetForecast with days parameter (1-7)
- **THEN** return forecast for specified number of days

### Requirement: Weather API Integration
The service SHALL integrate with a public weather API to fetch weather data.

#### Scenario: API returns valid response
- **WHEN** weather API returns 200 OK with valid JSON
- **THEN** parse and map response to domain models
- **THEN** return formatted weather data to client

#### Scenario: API rate limit exceeded
- **WHEN** weather API returns 429 Too Many Requests
- **THEN** return UNAVAILABLE error with retry-after hint
- **AND** serve cached data if available

#### Scenario: API unavailable
- **WHEN** weather API is unreachable or returns error
- **THEN** return UNAVAILABLE error
- **AND** log error for monitoring

### Requirement: Weather Data Caching
The service SHALL cache weather data to reduce API calls.

#### Scenario: Cache hit
- **WHEN** request matches cached data within TTL
- **THEN** return cached response immediately

#### Scenario: Cache miss
- **WHEN** no valid cached data exists
- **THEN** fetch from API and cache result

#### Scenario: Cache expiration
- **WHEN** cached data exceeds TTL
- **THEN** fetch fresh data from API
- **AND** update cache with new data

### Requirement: Configuration
The service SHALL support configuration via environment variables.

#### Scenario: API key configured
- **WHEN** WEATHER_API_KEY environment variable is set
- **THEN** use configured key for API requests

#### Scenario: Missing API key
- **WHEN** WEATHER_API_KEY is not configured
- **THEN** return error at service startup
- **AND** log warning about missing configuration
