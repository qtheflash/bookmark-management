package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qtheflash/bookmark-management/internal/service"
)

type CheckHealth interface {
	CheckHealthImpl(c *gin.Context)
}

type checkHealthHandler struct {
	checkHealthService service.CheckHealth
}

func NewCheckHealth(checkHealthSvc service.CheckHealth) CheckHealth {
	return &checkHealthHandler{
		checkHealthService: checkHealthSvc,
	}
}

func (h *checkHealthHandler) CheckHealthImpl(c *gin.Context) {
	resp := h.checkHealthService.CheckHealthImpl()
	c.JSON(http.StatusOK, resp)
}
