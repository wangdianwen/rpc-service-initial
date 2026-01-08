package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"rpc-service/internal/infrastructure/configuration"
	"rpc-service/internal/infrastructure/di"
	"rpc-service/internal/infrastructure/logging"
	rpchandler "rpc-service/internal/interface/rpc/service"
)

func main() {
	logger := logging.NewSimpleLogger()
	logger.Info("Starting RPC Service...")

	appService, err := di.NewApplicationService()
	if err != nil {
		logger.Error("Failed to create application service: %v", err)
		log.Fatalf("Failed to create application service: %v", err)
	}

	rpcHandler := rpchandler.NewServiceHandler(appService, logger)

	configProvider := configuration.NewEnvConfigProvider()
	config, err := configProvider.GetConfig()
	if err != nil {
		logger.Error("Failed to load config: %v", err)
		log.Fatalf("Failed to load config: %v", err)
	}

	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)

	rpc.Register(rpcHandler) //nolint:errcheck

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Error("Failed to listen: %v", err)
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()

	logger.Info("RPC Service listening on %s", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error("Failed to accept connection: %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
