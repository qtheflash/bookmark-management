package service

import (
	"github.com/google/uuid"
	"github.com/qtheflash/bookmark-management/internal/config"
)
// HealthCheckServiceResponse represents the response payload containing status information, service name, and instance identification for a health check request.
type HealthCheckServiceResponse struct {
	Message		string	`json:"message"`
	ServiceName	string	`json:"service_name"`
	InstanceID	string	`json:"instance_id"`
}
// CheckHealth defines an interface for checking the health status of a service.
type CheckHealth interface {
	CheckHealthImpl() HealthCheckServiceResponse
}
// checkHealthService represents a service instance for health monitoring and check operations.
type checkHealthService struct {
	serviceName	string
	instanceID	string
}
// NewCheckHealth creates and initializes a new CheckHealth service with the provided configuration.
// If InstanceID is empty in the configuration, a new UUID is automatically generated.
func NewCheckHealth(cfg *config.Config) CheckHealth {
	serviceName := cfg.ServiceName
	instanceID := cfg.InstanceID
	if instanceID == "" {
		instanceID = uuid.New().String()
	}

	return &checkHealthService{
		serviceName:	serviceName,
		instanceID:	instanceID,
	}
}
// CheckHealthImpl performs the health check for the service instance and returns a HealthCheckServiceResponse containing the service details and health status.
func (h *checkHealthService) CheckHealthImpl() HealthCheckServiceResponse {
	return HealthCheckServiceResponse{
		Message:	"OK",
		ServiceName:	h.serviceName,
		InstanceID:	h.instanceID,
	}
}
