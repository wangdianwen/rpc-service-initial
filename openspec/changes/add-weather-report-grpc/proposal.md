# Change: Add Weather Report gRPC Service

## Why
The RPC service currently lacks weather reporting capabilities. Adding a weather report gRPC service will enable inter-service communication for weather data, supporting current conditions and multi-day forecasts sourced from a public weather API.

## What Changes
- Add new weather domain entities (WeatherReport, Forecast, Location)
- Create gRPC service definition in protobuf
- Implement weather repository and service layer
- Add weather gRPC handler
- Integrate with public weather API (OpenWeatherMap/WeatherAPI)
- Wire dependency injection for new components
- Add configuration for API key and endpoints

## Impact
- Affected specs: weather-report-rpc (new capability)
- Affected code:
  - `proto/weather.proto` (new gRPC definition)
  - `internal/domain/entity/weather*.go` (new domain models)
  - `internal/domain/repository/weather_repository.go` (new repository interface)
  - `internal/infrastructure/external/weather_api.go` (external API client)
  - `internal/application/service/weather_app.go` (application service)
  - `internal/interface/grpc/weather/*.go` (gRPC handler)
  - `internal/infrastructure/di/wire.go` (DI updates)
  - `cmd/rpc-service/main.go` (optional separate entrypoint)
- New dependencies: weather API client library
