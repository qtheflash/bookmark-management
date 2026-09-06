package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {

	type HealthCheckResponse struct {
		Message     string `json:"message"`
		ServiceName string `json:"service_name"`
		InstanceID  string `json:"instance_id"`
	}

	instanceID := ""
	if instanceID == "" {
		instanceID = uuid.New().String()
	}

	serviceName := "bookmark-sevice"

	r := gin.Default()
	r.GET("/health-check", func(c *gin.Context) {
		resp := HealthCheckResponse{
			Message:     "OK",
			ServiceName: serviceName,
			InstanceID:  instanceID,
		}
		c.JSON(http.StatusOK, resp)
	})

	r.Run(":8080")

}
