## Context
Add weather report capabilities to the existing RPC service using gRPC protocol. The service will fetch data from a public weather API (OpenWeatherMap free tier or similar) and expose it via gRPC for inter-service communication.

Constraints:
- Must follow existing DDD architecture patterns
- Must use Wire for dependency injection
- Must support both current weather and multi-day forecasts
- Must handle API rate limits gracefully
- Must provide mock implementation for testing

## Goals / Non-Goals

### Goals:
- Expose weather data via gRPC endpoints
- Support current weather and 5-day forecast queries
- Integrate with public weather API (OpenWeatherMap)
- Provide testable architecture with mock data source
- Follow existing project conventions

### Non-Goals:
- Historical weather data (out of scope for initial version)
- Weather alerts and notifications
- WebSocket for real-time updates
- Complex caching strategies (basic in-memory allowed)

## Decisions

### 1. Protocol: gRPC
- Chosen over standard net/rpc for better performance and generated clients
- Protobuf provides strong typing and backward compatibility
- Existing codebase already supports both HTTP and RPC patterns

### 2. Weather API: OpenWeatherMap
- Free tier available (60 calls/minute, 1000/day)
- Good Go client libraries available
- Supports current weather and 5-day forecast
- Alternative: WeatherAPI.com (also free tier)

### 3. Error Handling Strategy
- Return gRPC status codes for API errors
- Map weather API errors to domain errors
- Implement retry logic with exponential backoff

### 4. Testing Strategy
- Interface-based API client for mocking
- Table-driven tests for service layer
- gRPC reflection for integration testing

## Risks / Trade-offs

| Risk | Mitigation |
|------|------------|
| API rate limits | Implement caching with configurable TTL |
| API downtime | Return cached data or graceful degradation |
| API response changes | Version the proto, validate responses |
| Performance | Add in-memory cache, async background refresh |

## Migration Plan
1. Create new packages without modifying existing code
2. Add new gRPC server alongside existing RPC server
3. Wire new components into existing DI
4. Deploy alongside existing service
5. Route weather requests to new service

## Open Questions
- Should weather endpoints be in separate gRPC server/port?
- Cache TTL configuration - default 15 minutes?
- Support for multiple weather API providers?
