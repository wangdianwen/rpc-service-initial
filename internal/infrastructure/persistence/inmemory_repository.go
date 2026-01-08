package persistence

import (
	"context"
	"errors"
	"sync"

	"rpc-service/internal/domain"
	"rpc-service/internal/domain/entity"
)

type InMemoryServiceRepository struct {
	mu        sync.RWMutex
	services  map[string]*entity.Service
	nameIndex map[string]string
}

func NewInMemoryServiceRepository() *InMemoryServiceRepository {
	return &InMemoryServiceRepository{
		services:  make(map[string]*entity.Service),
		nameIndex: make(map[string]string),
	}
}

func (r *InMemoryServiceRepository) FindByID(ctx context.Context, id string) (*entity.Service, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	service, exists := r.services[id]
	if !exists {
		return nil, domain.ErrServiceNotFound
	}

	return service, nil
}

func (r *InMemoryServiceRepository) FindAll(ctx context.Context) ([]*entity.Service, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	services := make([]*entity.Service, 0, len(r.services))
	for _, service := range r.services {
		services = append(services, service)
	}

	return services, nil
}

func (r *InMemoryServiceRepository) Save(ctx context.Context, service *entity.Service) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.services[service.ID]; exists {
		return errors.New("service already exists")
	}

	r.services[service.ID] = service
	r.nameIndex[service.Name] = service.ID

	return nil
}

func (r *InMemoryServiceRepository) Update(ctx context.Context, service *entity.Service) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	existing, exists := r.services[service.ID]
	if !exists {
		return domain.ErrServiceNotFound
	}

	if existing.Name != service.Name {
		delete(r.nameIndex, existing.Name)
		r.nameIndex[service.Name] = service.ID
	}

	r.services[service.ID] = service

	return nil
}

func (r *InMemoryServiceRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	service, exists := r.services[id]
	if !exists {
		return domain.ErrServiceNotFound
	}

	delete(r.services, id)
	delete(r.nameIndex, service.Name)

	return nil
}

func (r *InMemoryServiceRepository) FindByName(ctx context.Context, name string) (*entity.Service, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	id, exists := r.nameIndex[name]
	if !exists {
		return nil, domain.ErrServiceNotFound
	}

	return r.services[id], nil
}
