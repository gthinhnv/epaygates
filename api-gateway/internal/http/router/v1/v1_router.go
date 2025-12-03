package v1

import (
	"apigateway/internal/http/router/v1/staticpagerouter"
	"net/http"

	"github.com/gin-gonic/gin"
)

func New(router *gin.Engine) *gin.Engine {
	routerV1 := router.Group("/v1")

	routerV1.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	staticpagerouter.RegisterRoutes(routerV1)

	return router
}
