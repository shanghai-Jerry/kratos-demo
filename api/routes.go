package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shanghai-Jerry/krato-demo/internal/service"
)

const (
	BaseAPI = ""
)

// RegisterRoutes the function which register all routes
func RegisterRoutes(router *gin.Engine, service *service.DemoService, middlewareList ...gin.HandlerFunc) {
	// basic url group use basic middleware
	group := router.Group(BaseAPI, middlewareList...)
	group.GET("/healthCheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})
}
