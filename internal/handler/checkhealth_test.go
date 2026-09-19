package handler

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/qtheflash/bookmark-management/internal/service"
	"github.com/stretchr/testify/assert"
)

type checkHealthServiceStub struct {
	response service.HealthCheckServiceResponse
}

// CheckHealthImpl returns the configured health check response.
func (s *checkHealthServiceStub) CheckHealthImpl() service.HealthCheckServiceResponse {
	return s.response
}

// TestCheckHealthImpl verifies that the handler returns the service response as JSON.
func TestCheckHealthImpl(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)
	engine := gin.New()
	serviceResponse := service.HealthCheckServiceResponse{
		Message:     "OK",
		ServiceName: "bookmark-api",
		InstanceID:  "my-instance-id",
	}
	handler := NewCheckHealth(&checkHealthServiceStub{response: serviceResponse})
	engine.GET("/health-check", handler.CheckHealthImpl)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/health-check", nil)
	engine.ServeHTTP(recorder, request)

	assert.Equal(t, 200, recorder.Code)
	assert.JSONEq(t, `{"message":"OK","service_name":"bookmark-api","instance_id":"my-instance-id"}`, recorder.Body.String())
}
