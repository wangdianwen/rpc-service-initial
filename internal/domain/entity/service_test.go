package entity

import (
	"strings"
	"testing"
)

func TestNewService(t *testing.T) {
	name := "test-service"
	data := "test data"

	service, err := NewService(name, data)
	if err != nil {
		t.Fatalf("NewService returned error: %v", err)
	}

	if service.ID == "" {
		t.Error("Service ID should not be empty")
	}

	if service.Name != name {
		t.Errorf("Expected name %s, got %s", name, service.Name)
	}

	if service.Data != data {
		t.Errorf("Expected data %s, got %s", data, service.Data)
	}

	if service.CreatedAt.IsZero() {
		t.Error("CreatedAt should not be zero")
	}

	if service.UpdatedAt.IsZero() {
		t.Error("UpdatedAt should not be zero")
	}
}

func TestServiceValidation(t *testing.T) {
	tests := []struct {
		name        string
		serviceName string
		expectError bool
	}{
		{
			name:        "valid service",
			serviceName: "valid-name",
			expectError: false,
		},
		{
			name:        "empty name",
			serviceName: "",
			expectError: true,
		},
		{
			name:        "name too long",
			serviceName: strings.Repeat("a", 101),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service, err := NewService(tt.serviceName, "data")

			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
				if service != nil {
					t.Error("Expected nil service on error")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if service == nil {
					t.Error("Expected service but got nil")
				}
			}
		})
	}
}

func TestServiceUpdate(t *testing.T) {
	service, err := NewService("test", "initial data")
	if err != nil {
		t.Fatalf("NewService returned error: %v", err)
	}

	originalID := service.ID
	originalCreatedAt := service.CreatedAt

	service.Update("updated data")

	if service.ID != originalID {
		t.Error("ID should not change on update")
	}

	if service.Data != "updated data" {
		t.Errorf("Expected data 'updated data', got '%s'", service.Data)
	}

	if service.CreatedAt != originalCreatedAt {
		t.Error("CreatedAt should not change on update")
	}

	if !service.UpdatedAt.After(service.CreatedAt) {
		t.Error("UpdatedAt should be after CreatedAt after update")
	}
}

func TestServiceUpdateName(t *testing.T) {
	service, err := NewService("original", "data")
	if err != nil {
		t.Fatalf("NewService returned error: %v", err)
	}

	err = service.UpdateName("updated")
	if err != nil {
		t.Errorf("UpdateName returned error: %v", err)
	}

	if service.Name != "updated" {
		t.Errorf("Expected name 'updated', got '%s'", service.Name)
	}
}

func TestServiceUpdateNameValidation(t *testing.T) {
	service, err := NewService("original", "data")
	if err != nil {
		t.Fatalf("NewService returned error: %v", err)
	}

	err = service.UpdateName("")
	if err == nil {
		t.Error("Expected error for empty name")
	}
}
