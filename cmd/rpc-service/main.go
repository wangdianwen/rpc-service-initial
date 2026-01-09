package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"

	"google.golang.org/grpc"
	"rpc-service/internal/infrastructure/configuration"
	"rpc-service/internal/infrastructure/di"
	"rpc-service/internal/infrastructure/external"
	"rpc-service/internal/infrastructure/logging"
	weathergrpc "rpc-service/internal/interface/grpc/weather"
	rpchandler "rpc-service/internal/interface/rpc/service"
	pb "rpc-service/pb/weather"
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

	weatherConfig := external.WeatherAPIConfig{
		APIKey:  config.WeatherAPIKey,
		BaseURL: config.WeatherAPIURL,
		Timeout: time.Duration(config.WeatherTimeout) * time.Second,
	}
	weatherClient := external.NewWeatherAPIClient(weatherConfig)
	weatherRepo := external.NewWeatherRepositoryWithCache(weatherClient, 15*time.Minute)
	weatherService := di.NewWeatherService(weatherRepo)
	weatherServer := weathergrpc.NewWeatherServer(weatherService)

	grpcAddr := fmt.Sprintf("%s:%d", config.Host, config.Port+1)

	grpcServer := grpc.NewServer()
	pb.RegisterWeatherServiceServer(grpcServer, weatherServer)

	go func() {
		grpcListener, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			logger.Error("Failed to listen for gRPC: %v", err)
			return
		}
		logger.Info("Weather gRPC Service listening on %s", grpcAddr)
		if err := grpcServer.Serve(grpcListener); err != nil {
			logger.Error("gRPC server error: %v", err)
		}
	}()

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
