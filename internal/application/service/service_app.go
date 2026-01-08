package service

import (
	"context"
	"errors"

	"rpc-service/internal/application/dto"
	"rpc-service/internal/application/mapper"
	"rpc-service/internal/domain"
	"rpc-service/internal/domain/entity"
	"rpc-service/internal/domain/repository"
)

type ApplicationService struct {
	repo repository.ServiceRepository
}

func NewApplicationService(repo repository.ServiceRepository) *ApplicationService {
	return &ApplicationService{repo: repo}
}

func (s *ApplicationService) GetService(ctx context.Context, id string) (*dto.ServiceResponse, error) {
	service, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrServiceNotFound) {
			return nil, &domain.NotFoundError{Resource: "service", ID: id}
		}
		return nil, err
	}
	return mapper.ToResponse(service), nil
}

func (s *ApplicationService) GetAllServices(ctx context.Context) ([]*dto.ServiceResponse, error) {
	services, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.ToResponses(services), nil
}

func (s *ApplicationService) CreateService(ctx context.Context, req *dto.CreateServiceRequest) (*dto.ServiceResponse, error) {
	newEntity, err := entity.NewService(req.Name, req.Data)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Save(ctx, newEntity); err != nil {
		return nil, err
	}

	return mapper.ToResponse(newEntity), nil
}

func (s *ApplicationService) UpdateService(ctx context.Context, id string, req *dto.UpdateServiceRequest) (*dto.ServiceResponse, error) {
	service, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrServiceNotFound) {
			return nil, &domain.NotFoundError{Resource: "service", ID: id}
		}
		return nil, err
	}

	if req.Name != "" {
		if err := service.UpdateName(req.Name); err != nil {
			return nil, err
		}
	}

	if req.Data != "" {
		service.Update(req.Data)
	}

	if err := s.repo.Update(ctx, service); err != nil {
		return nil, err
	}

	return mapper.ToResponse(service), nil
}

func (s *ApplicationService) DeleteService(ctx context.Context, id string) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrServiceNotFound) {
			return &domain.NotFoundError{Resource: "service", ID: id}
		}
	}
	return err
}
