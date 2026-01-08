package service

import (
	"context"
	"encoding/json"

	"rpc-service/internal/application/dto"
	"rpc-service/internal/application/service"
	"rpc-service/internal/infrastructure/logging"
)

type ServiceHandler struct {
	*service.ApplicationService
	logger *logging.SimpleLogger
}

func NewServiceHandler(appService *service.ApplicationService, logger *logging.SimpleLogger) *ServiceHandler {
	return &ServiceHandler{
		ApplicationService: appService,
		logger:             logger,
	}
}

type RPCRequest struct {
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
}

type RPCResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func (h *ServiceHandler) HandleRequest(ctx context.Context, req *RPCRequest) *RPCResponse {
	switch req.Method {
	case "GetService":
		return h.handleGetService(ctx, req)
	case "GetAllServices":
		return h.handleGetAllServices(ctx, req)
	case "CreateService":
		return h.handleCreateService(ctx, req)
	case "UpdateService":
		return h.handleUpdateService(ctx, req)
	case "DeleteService":
		return h.handleDeleteService(ctx, req)
	default:
		return &RPCResponse{
			Success: false,
			Error:   map[string]string{"message": "unknown method: " + req.Method},
		}
	}
}

func (h *ServiceHandler) handleGetService(ctx context.Context, req *RPCRequest) *RPCResponse {
	var params struct {
		ID string `json:"id"`
	}

	if err := json.Unmarshal(req.Params, &params); err != nil {
		return &RPCResponse{
			Success: false,
			Error:   map[string]string{"message": "invalid params"},
		}
	}

	result, err := h.GetService(ctx, params.ID)
	if err != nil {
		h.logger.Error("GetService failed: %v", err)
		return &RPCResponse{
			Success: false,
			Error:   map[string]string{"message": err.Error()},
		}
	}

	return &RPCResponse{
		Success: true,
		Data:    result,
	}
}

func (h *ServiceHandler) handleGetAllServices(ctx context.Context, req *RPCRequest) *RPCResponse {
	results, err := h.GetAllServices(ctx)
	if err != nil {
		h.logger.Error("GetAllServices failed: %v", err)
		return &RPCResponse{
			Success: false,
			Error:   map[string]string{"message": err.Error()},
		}
	}

	return &RPCResponse{
		Success: true,
		Data:    results,
	}
}

func (h *ServiceHandler) handleCreateService(ctx context.Context, req *RPCRequest) *RPCResponse {
	var params dto.CreateServiceRequest

	if err := json.Unmarshal(req.Params, &params); err != nil {
		return &RPCResponse{
			Success: false,
			Error:   map[string]string{"message": "invalid params"},
		}
	}

	result, err := h.CreateService(ctx, &params)
	if err != nil {
		h.logger.Error("CreateService failed: %v", err)
		return &RPCResponse{
			Success: false,
			Error:   map[string]string{"message": err.Error()},
		}
	}

	return &RPCResponse{
		Success: true,
		Data:    result,
	}
}

func (h *ServiceHandler) handleUpdateService(ctx context.Context, req *RPCRequest) *RPCResponse {
	var params struct {
		ID   string                   `json:"id"`
		Data dto.UpdateServiceRequest `json:"data"`
	}

	if err := json.Unmarshal(req.Params, &params); err != nil {
		return &RPCResponse{
			Success: false,
			Error:   map[string]string{"message": "invalid params"},
		}
	}

	result, err := h.UpdateService(ctx, params.ID, &params.Data)
	if err != nil {
		h.logger.Error("UpdateService failed: %v", err)
		return &RPCResponse{
			Success: false,
			Error:   map[string]string{"message": err.Error()},
		}
	}

	return &RPCResponse{
		Success: true,
		Data:    result,
	}
}

func (h *ServiceHandler) handleDeleteService(ctx context.Context, req *RPCRequest) *RPCResponse {
	var params struct {
		ID string `json:"id"`
	}

	if err := json.Unmarshal(req.Params, &params); err != nil {
		return &RPCResponse{
			Success: false,
			Error:   map[string]string{"message": "invalid params"},
		}
	}

	err := h.DeleteService(ctx, params.ID)
	if err != nil {
		h.logger.Error("DeleteService failed: %v", err)
		return &RPCResponse{
			Success: false,
			Error:   map[string]string{"message": err.Error()},
		}
	}

	return &RPCResponse{
		Success: true,
	}
}
