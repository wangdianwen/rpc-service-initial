# Project Context

## Purpose
Go-based microservice providing API and RPC endpoints for inter-service communication. Uses standard library `net/http` for HTTP APIs and RPC for service-to-service calls. Stateless service with Wire dependency injection.

## Tech Stack
- **Language**: Go
- **HTTP**: Standard library `net/http`
- **RPC**: Go built-in RPC / gRPC
- **Dependency Injection**: Wire
- **Architecture**: Domain-driven design

## Project Conventions

### Code Style
- Follow Go standard formatting (`go fmt`)
- Use `golint`/`staticcheck` for linting
- Package naming: lowercase, short, descriptive (e.g., `handler`, `service`, `repository`)
- Struct tags for serialization: `json:"field_name"` and `yaml:"field_name"`

### Architecture Patterns
- Domain-driven design with clear boundaries between layers
- Separation of concerns: handlers → services → domain entities
- Dependency injection via Wire for testability
- Stateless service design

### Testing Strategy
- Unit tests with standard library `testing` package
- Table-driven tests preferred for function testing
- Integration tests for handlers/services
- Use Wire-generated providers for mock injection in tests

### Git Workflow
- **Branching**: GitFlow (main → develop → feature branches)
- **Feature branches**: `feature/description-name`
- **Bugfix branches**: `bugfix/description-name`
- **Release branches**: `release/v*.*.*`
- **Commit messages**: Conventional commits (`feat:`, `fix:`, `docs:`, `refactor:`)

## Domain Context
- Microservice architecture with HTTP and RPC interfaces
- Stateless operations - no local database
- Inter-service communication patterns

## Important Constraints
- Stateless service design
- Must support both HTTP and RPC protocols
- Dependency injection required via Wire

## External Dependencies
- Wire: Dependency injection compiler
- Standard library only for HTTP/RPC (no external frameworks)
