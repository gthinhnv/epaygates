package staticpagerouter

import (
	"apigateway/internal/http/handlers/v1/staticpagehandler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	staticPageHandler := staticpagehandler.NewStaticPageHandler()

	router.POST("/staticPages/create", staticPageHandler.Create)
	router.GET("/staticPages/:id", staticPageHandler.Get)
	router.GET("/staticPages", staticPageHandler.List)
}
