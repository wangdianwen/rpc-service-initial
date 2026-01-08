.PHONY: all clean proto run build

all: proto build

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*.proto

build:
	go build -o bin/rpc-service main.go

run: build
	./bin/rpc-service

clean:
	rm -rf bin/
