## 1. Domain Layer
- [x] 1.1 Create `internal/domain/entity/weather.go` with WeatherReport, Forecast, Location entities
- [x] 1.2 Create `internal/domain/error.go` extension for weather-specific errors
- [x] 1.3 Create `internal/domain/repository/weather_repository.go` interface

## 2. Infrastructure Layer
- [x] 2.1 Create `internal/infrastructure/external/weather_api.go` for public API client
- [x] 2.2 Add weather API configuration to `internal/infrastructure/configuration/config.go`
- [x] 2.3 Update `internal/infrastructure/di/factory.go` for Wire providers

## 3. Application Layer
- [x] 3.1 Create `internal/application/dto/weather_dto.go` for request/response DTOs
- [x] 3.2 Create `internal/application/mapper/weather_mapper.go`
- [x] 3.3 Create `internal/application/validator/weather_validator.go`
- [x] 3.4 Create `internal/application/service/weather_app.go` with GetCurrentWeather, GetForecast methods

## 4. Interface Layer (gRPC)
- [x] 4.1 Create `proto/weather.proto` with GetCurrentWeather and GetForecast RPC methods
- [x] 4.2 Generate Go code from protobuf (`protoc` generation)
- [x] 4.3 Create `internal/interface/grpc/weather/handler.go` gRPC handler
- [x] 4.4 Register weather gRPC service in main

## 5. Configuration
- [x] 5.1 Add WEATHER_API_KEY, WEATHER_API_URL to environment configuration
- [x] 5.2 Update config loading in `internal/infrastructure/configuration/`

## 6. Testing
- [x] 6.1 Add unit tests for weather entities
- [x] 6.2 Add unit tests for weather service
- [ ] 6.3 Add integration tests for gRPC handlers
- [x] 6.4 Add mock weather API client for tests

## 7. Documentation
- [x] 7.1 Update README with weather gRPC usage examples
- [x] 7.2 Add proto documentation comments
