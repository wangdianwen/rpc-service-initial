package mapper

import (
	"time"

	"rpc-service/internal/application/dto"
	"rpc-service/internal/domain/entity"
)

func ToResponse(service *entity.Service) *dto.ServiceResponse {
	return &dto.ServiceResponse{
		ID:        service.ID,
		Name:      service.Name,
		Data:      service.Data,
		CreatedAt: service.CreatedAt.Format(time.RFC3339),
		UpdatedAt: service.UpdatedAt.Format(time.RFC3339),
	}
}

func ToResponses(services []*entity.Service) []*dto.ServiceResponse {
	responses := make([]*dto.ServiceResponse, len(services))
	for i, service := range services {
		responses[i] = ToResponse(service)
	}
	return responses
}

func ToEntity(req *dto.ServiceRequest) (*entity.Service, error) {
	return entity.NewService(req.Name, req.Data)
}
