# RPC Service

![Go Version](https://img.shields.io/badge/Go-1.22-blue)
![License](https://img.shields.io/badge/License-MIT-green)
![Build Status](https://github.com/wangdianwen/rpc-service-initial/actions/workflows/go.yml/badge.svg?branch=main)

A Go-based RPC (Remote Procedure Call) service built with Domain-Driven Design (DDD) architecture.

## Features

- RPC server with JSON-RPC support
- Domain-Driven Design architecture
- Dependency injection
- Configuration management via environment variables
- Structured logging

## Getting Started

### Prerequisites

- Go 1.22 or higher
- Make

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd rpc-service
```

2. Install dependencies:
```bash
make install-tools
```

3. Build the service:
```bash
make build
```

4. Run the service:
```bash
make run
```

## Development

### Available Make Targets

- `make build` - Build the binary
- `make run` - Build and run the service
- `make test` - Run tests with race detection
- `make test-coverage` - Run tests with coverage report
- `make lint` - Run golangci-lint
- `make format` - Format code with gofumpt and goimports
- `make check` - Run format, lint, and test
- `make proto` - Generate protocol buffer code
- `make clean` - Remove build artifacts
- `make install-tools` - Install development tools

### Project Structure

```
.
├── cmd/
│   └── rpc-service/
│       └── main.go           # Application entry point
├── internal/
│   ├── domain/               # Domain layer (entities, value objects, etc.)
│   ├── application/          # Application services
│   ├── infrastructure/       # Infrastructure (config, DI, logging)
│   └── interface/            # Interface adapters (RPC handlers, etc.)
├── proto/
│   └── service.proto         # Protocol buffer definitions
├── .github/workflows/        # CI/CD pipelines
├── .golangci.yml             # Linter configuration
├── Makefile                  # Build targets
└── go.mod                    # Go module definition
```

## Testing

Run all tests:
```bash
make test
```

Run tests with coverage report:
```bash
make test-coverage
```

## License

MIT License
