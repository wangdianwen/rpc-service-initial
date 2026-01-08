package dto

type ServiceRequest struct {
	Name string `json:"name" validate:"required,max=100"`
	Data string `json:"data"`
}

type ServiceResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Data      string `json:"data"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateServiceRequest struct {
	Name string `json:"name" validate:"required,max=100"`
	Data string `json:"data"`
}

type UpdateServiceRequest struct {
	Name string `json:"name" validate:"max=100"`
	Data string `json:"data"`
}
