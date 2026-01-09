package di

import (
	"rpc-service/internal/application/service"
	"rpc-service/internal/domain/repository"
	"rpc-service/internal/infrastructure/persistence"
)

func NewApplicationService() (*service.ApplicationService, error) {
	repo := persistence.NewInMemoryServiceRepository()
	return service.NewApplicationService(repo), nil
}

func NewWeatherService(repo repository.WeatherRepository) *service.WeatherService {
	return service.NewWeatherService(repo)
}
