package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/qtheflash/bookmark-management/internal/config"
	"github.com/qtheflash/bookmark-management/internal/handler"
	"github.com/qtheflash/bookmark-management/internal/service"
)

type Engine interface {
	Start() error
}

type engine struct {
	app *gin.Engine
	cfg *config.Config
}

func NewEngine(cfg *config.Config) Engine {
	app := &engine{
		app: gin.Default(),
		cfg: cfg,
	}
	app.initRoutes()
	return app
}

func (e *engine) Start() error {
	return e.app.Run(fmt.Sprintf(":%s", e.cfg.AppPort))
}

func (e *engine) initRoutes() {
	checkHealthSvc := service.NewCheckHealth(e.cfg)
	checkHealthHandler := handler.NewCheckHealth(checkHealthSvc)

	e.app.GET("/health-check", checkHealthHandler.CheckHealthImpl)
}
