package entity

import (
	"fmt"
	"time"

	"rpc-service/internal/domain"
)

const (
	MaxServiceNameLength = 100
)

type Service struct {
	ID        string
	Name      string
	Data      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewService(name, data string) (*Service, error) {
	service := &Service{
		ID:        generateID(),
		Name:      name,
		Data:      data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := service.Validate(); err != nil {
		return nil, err
	}

	return service, nil
}

func generateID() string {
	return fmt.Sprintf("%d-%d", time.Now().UnixNano(), time.Now().Unix())
}

func (s *Service) Validate() error {
	if s.Name == "" {
		return &domain.DomainError{
			Code:    "INVALID_NAME",
			Message: "service name cannot be empty",
			Err:     domain.ErrInvalidServiceName,
		}
	}

	if len(s.Name) > MaxServiceNameLength {
		return &domain.DomainError{
			Code:    "NAME_TOO_LONG",
			Message: fmt.Sprintf("service name exceeds maximum length of %d", MaxServiceNameLength),
			Err:     domain.ErrServiceNameTooLong,
		}
	}

	return nil
}

func (s *Service) Update(data string) {
	s.Data = data
	s.UpdatedAt = time.Now()
}

func (s *Service) UpdateName(name string) error {
	s.Name = name
	s.UpdatedAt = time.Now()
	return s.Validate()
}
