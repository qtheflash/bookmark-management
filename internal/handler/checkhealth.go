package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qtheflash/bookmark-management/internal/service"
)
// CheckHealth defines the interface for handling application health check requests.
type CheckHealth interface {
	CheckHealthImpl(c *gin.Context)
}
// checkHealthHandler handles health check requests using the underlying health check service.
type checkHealthHandler struct {
	checkHealthService service.CheckHealth
}
// NewCheckHealth creates and returns a new CheckHealth handler instance initialized with the given service.
func NewCheckHealth(checkHealthSvc service.CheckHealth) CheckHealth {
	return &checkHealthHandler{
		checkHealthService: checkHealthSvc,
	}
}
// CheckHealthImpl handles the health check HTTP request and returns the status response as JSON.
func (h *checkHealthHandler) CheckHealthImpl(c *gin.Context) {
	resp := h.checkHealthService.CheckHealthImpl()
	c.JSON(http.StatusOK, resp)
}
