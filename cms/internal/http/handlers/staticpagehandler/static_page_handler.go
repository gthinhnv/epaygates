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

func (ctrl *StaticPageHandler) Create(c *gin.Context) {
	p := staticpage.Create{
		BasePage: layout.BasePage{Context: c, Lang: c.GetString("lang")},
	}
	layout.WritePageTemplate(c.Writer, &p)
}

func (h *StaticPageHandler) List(c *gin.Context) {
	p := staticpage.List{
		BasePage: layout.BasePage{Context: c, Lang: c.GetString("lang")},
	}
	layout.WritePageTemplate(c.Writer, &p)
}
