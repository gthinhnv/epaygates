package v1

import (
	"cms/internal/http/handler/v1/staticpagehandler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func New(router *gin.Engine) *gin.Engine {
	routerV1 := router.Group("/v1")

	routerV1.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	staticPageHandler := staticpagehandler.NewStaticPageHandler()
	staticPageHandler.RegisterRoutes(routerV1)

	return router
}
