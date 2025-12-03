package commonrouter

import (
	"apigateway/internal/http/handlers/v1/commonhandler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	commonHandler := commonhandler.NewCommonHandler()

	router.GET("/common/statuses", commonHandler.ListStatuses)
}
