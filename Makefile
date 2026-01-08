.PHONY: all clean proto build run test test-coverage lint format check install-tools generate mocks docker-build docker-run \
        help

all: proto build test lint

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
           --go-grpc_out=. --go-grpc_opt=paths=source_relative \
           proto/*.proto

build:
	go build -o bin/rpc-service ./cmd/rpc-service

run: build
	./bin/rpc-service

clean:
	rm -rf bin/
	rm -f coverage.out coverage.html
	rm -f *.coverprofile

test:
	go test -v -race -cover -count=1 ./...

test-coverage:
	go test -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

lint:
	golangci-lint run ./...

format:
	gofumpt -l -w .
	goimports -l -w -local rpc-service .

check: format lint test

install-tools:
	go install mvdan.cc/gofumpt@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/golang/mock/mockgen@latest
	go install gotest.tools/gotestsum@latest

generate:
	go generate ./...

docker-build:
	docker build -t rpc-service:latest .

docker-run:
	docker run -p 1234:1234 rpc-service:latest

help:
	@echo "Available targets:"
	@echo "  all              - Run proto, build, test, and lint"
	@echo "  build            - Build the binary"
	@echo "  run              - Build and run the service"
	@echo "  clean            - Remove build artifacts"
	@echo "  proto            - Generate protocol buffers"
	@echo "  test             - Run tests with race detection"
	@echo "  test-coverage    - Run tests with coverage report"
	@echo "  lint             - Run golangci-lint"
	@echo "  format           - Format code with gofumpt and goimports"
	@echo "  check            - Run format, lint, and test"
	@echo "  install-tools    - Install development tools"
	@echo "  generate         - Generate code (mocks, etc.)"
	@echo "  docker-build     - Build Docker image"
	@echo "  docker-run       - Run Docker container"
	@echo "  help             - Show this help"
