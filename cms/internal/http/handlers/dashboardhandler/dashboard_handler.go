package dashboardhandler

import (
	"cms/internal/http/views/layout"
	"cms/internal/http/views/pages/dashboard"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
}

func NewStaticPageHandler() *DashboardHandler {
	return &DashboardHandler{}
}

func (h *DashboardHandler) GetIndex(c *gin.Context) {
	p := dashboard.DashboardIndex{}
	p.Context = c
	layout.WritePageTemplate(c.Writer, &p)
}
