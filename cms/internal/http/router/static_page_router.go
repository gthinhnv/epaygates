package router

import (
	"cms/internal/http/handlers/staticpagehandler"

	"github.com/gin-gonic/gin"
)

func RegisterStaticPageRoutes(router *gin.RouterGroup) {
	staticPageHandler := staticpagehandler.NewStaticPageHandler()

	router.GET("/staticPages", staticPageHandler.List)
}
