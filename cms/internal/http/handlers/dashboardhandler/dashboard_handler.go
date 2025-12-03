package dashboardhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
}

func NewStaticPageHandler() *DashboardHandler {
	return &DashboardHandler{}
}

func (h *DashboardHandler) GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Dashboard Index",
	})
}
