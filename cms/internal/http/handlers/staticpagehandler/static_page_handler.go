package staticpagehandler

import (
	"cms/internal/http/views/layout"
	"cms/internal/http/views/pages/staticpage"

	"github.com/gin-gonic/gin"
)

type StaticPageHandler struct {
}

func NewStaticPageHandler() *StaticPageHandler {
	return &StaticPageHandler{}
}

func (h *StaticPageHandler) List(c *gin.Context) {
	p := staticpage.List{}
	p.Ctx = c
	layout.WritePageTemplate(c.Writer, &p)
}
