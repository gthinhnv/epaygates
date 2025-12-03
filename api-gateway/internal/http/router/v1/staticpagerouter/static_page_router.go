package staticpagerouter

import (
	"apigateway/internal/http/handlers/v1/staticpagehandler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	staticPageHandler := staticpagehandler.NewStaticPageHandler()
	router.GET("/staticPages/:id", staticPageHandler.Get)
}
