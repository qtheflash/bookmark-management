package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qtheflash/bookmark-management/internal/config"
	"github.com/qtheflash/bookmark-management/internal/handler"
	"github.com/qtheflash/bookmark-management/internal/service"
)

// Engine represents a runnable component or service that can be started.
type Engine interface {
	Start() error
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// engine represents the core application server structure, holding the Gin engine instance and configuration settings.
type engine struct {
	app *gin.Engine
	cfg *config.Config
}

// ServeHTTP allows the engine to handle HTTP requests through its Gin router.
func (e *engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.app.ServeHTTP(w, r)
}

// NewEngine creates and initializes a new Engine instance with the provided configuration.
// It sets up the default Gin router, initializes the application routes, and returns the engine.
func NewEngine(cfg *config.Config) Engine {
	app := &engine{
		app: gin.Default(),
		cfg: cfg,
	}
	app.initRoutes()
	return app
}

// Start starts the application server on the configured application port.
// It returns an error if the application fails to start or run.
func (e *engine) Start() error {
	return e.app.Run(fmt.Sprintf(":%s", e.cfg.AppPort))
}

// initRoutes initializes the HTTP routes and registers the corresponding handlers for the engine.
func (e *engine) initRoutes() {
	checkHealthSvc := service.NewCheckHealth(e.cfg)
	checkHealthHandler := handler.NewCheckHealth(checkHealthSvc)

	e.app.GET("/health-check", checkHealthHandler.CheckHealthImpl)
}
