package staticpagehandler

import (
	"apigateway/gen/go/staticpagepb"
	"context"
	"net/http"
	"shared/models/staticpagemodel"
	"shared/pkg/utils/dbutil"
	"strconv"
	"web/internal/bootstrap"
	"web/internal/http/views/layout"
	"web/internal/http/views/pages/staticpage"

	"github.com/gin-gonic/gin"
)

type StaticPageHandler struct {
	client staticpagepb.StaticPageServiceClient
}

func NewStaticPageHandler() *StaticPageHandler {
	return &StaticPageHandler{
		client: staticpagepb.NewStaticPageServiceClient(bootstrap.APIServiceConn),
	}
}

func (h *StaticPageHandler) Create(c *gin.Context) {
	p := staticpage.Create{
		BasePage: layout.BasePage{Context: c},
	}
	layout.WritePageTemplate(c.Writer, &p)
}

func (h *StaticPageHandler) Update(c *gin.Context) {
	// Parse ID
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	resp, err := h.client.Get(context.Background(), &staticpagepb.GetRequest{
		Id: id,
	})
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var staticPage staticpagemodel.StaticPage
	err = dbutil.MapStruct(resp.Page, &staticPage)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	p := staticpage.Update{
		BasePage:   layout.BasePage{Context: c},
		StaticPage: &staticPage,
	}
	layout.WritePageTemplate(c.Writer, &p)
}

func (h *StaticPageHandler) List(c *gin.Context) {
	p := staticpage.List{
		BasePage: layout.BasePage{Context: c},
	}
	layout.WritePageTemplate(c.Writer, &p)
}
