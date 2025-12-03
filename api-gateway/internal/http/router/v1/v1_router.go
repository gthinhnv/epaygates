package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func New(router *gin.Engine) *gin.Engine {
	routerV1 := router.Group("/v1")

	routerV1.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	RegisterStaticPageRoutes(routerV1)

	return router
}
