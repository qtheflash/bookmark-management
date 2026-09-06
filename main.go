package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	type HealthCheckResponse struct {
		Message     string
		ServiceName string
		InstanceID  string
	}

	serviceName := "bookmark-sevice"
	instanceID := "abcdef12345"

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
