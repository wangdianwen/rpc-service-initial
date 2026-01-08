package repository

import (
	"context"

	"rpc-service/internal/domain/entity"
)

type ServiceRepository interface {
	FindByID(ctx context.Context, id string) (*entity.Service, error)
	FindAll(ctx context.Context) ([]*entity.Service, error)
	Save(ctx context.Context, service *entity.Service) error
	Update(ctx context.Context, service *entity.Service) error
	Delete(ctx context.Context, id string) error
	FindByName(ctx context.Context, name string) (*entity.Service, error)
}
