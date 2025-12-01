package staticpagehandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StaticPageHandler struct {
}

func NewStaticPageHandler() *StaticPageHandler {
	return &StaticPageHandler{}
}

func (h *StaticPageHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/staticPages/:id", h.GetPage)
}

func (h *StaticPageHandler) GetPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}
