package service

import (
	"github.com/google/uuid"
	"github.com/qtheflash/bookmark-management/internal/config"
)

type HealthCheckServiceResponse struct {
	Message     string `json:"message"`
	ServiceName string `json:"service_name"`
	InstanceID  string `json:"instance_id"`
}

type CheckHealth interface {
	CheckHealthImpl() HealthCheckServiceResponse
}

type checkHealthService struct {
	serviceName string
	instanceID  string
}

func NewCheckHealth(cfg *config.Config) CheckHealth {
	serviceName := cfg.ServiceName
	instanceID := ""
	if instanceID == "" {
		instanceID = uuid.New().String()
	}

	return &checkHealthService{
		serviceName: serviceName,
		instanceID:  instanceID,
	}
}

func (h *checkHealthService) CheckHealthImpl() HealthCheckServiceResponse {
	return HealthCheckServiceResponse{
		Message:     "OK",
		ServiceName: h.serviceName,
		InstanceID:  h.instanceID,
	}
}
