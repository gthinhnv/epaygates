package staticpagehandler

import (
	"cms/gen/go/staticpagepb"
	"cms/internal/bootstrap"
	"context"
	"net/http"
	"shared/models/staticpage"
	"shared/pkg/utils/apiutil"
	"shared/pkg/utils/dbutil"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type StaticPageHandler struct {
	client staticpagepb.StaticPageServiceClient
}

func NewStaticPageHandler() *StaticPageHandler {
	return &StaticPageHandler{
		client: staticpagepb.NewStaticPageServiceClient(bootstrap.MetadataServiceConn),
	}
}

func (h *StaticPageHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/staticPages/:id", h.Get)
}

func (h *StaticPageHandler) Get(c *gin.Context) {
	// Parse ID
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiutil.Response{
			Code:    apiutil.CODE_NOT_FOUND,
			Message: "Record doen't exist",
		})
		return
	}

	resp, err := h.client.Get(context.Background(), &staticpagepb.GetRequest{
		Id: id,
	})
	if err != nil {
		bootstrap.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Warn("StaticPageHandler >>> Get: failed to get static page")

		c.JSON(http.StatusBadRequest, apiutil.Response{
			Code:    apiutil.CODE_ERROR,
			Message: "An error happened when getting data",
		})
		return
	}

	var staticPage staticpage.StaticPage
	err = dbutil.MapStruct(resp.Page, &staticPage)
	if err != nil {
		bootstrap.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Warn("StaticPageHandler >>> Get: failed to get map struct proto to model")

		c.JSON(http.StatusBadRequest, apiutil.Response{
			Code:    apiutil.CODE_ERROR,
			Message: "An error happened when getting data",
		})
		return
	}

	c.JSON(http.StatusOK, apiutil.Response{
		Code:    apiutil.CODE_SUCCESS,
		Message: "Success",
		Data:    staticPage,
	})
}
