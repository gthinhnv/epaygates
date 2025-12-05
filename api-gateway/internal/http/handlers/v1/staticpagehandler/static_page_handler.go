package staticpagehandler

import (
	"apigateway/gen/go/staticpagepb"
	"apigateway/internal/bootstrap"
	"context"
	"fmt"
	"net/http"
	"shared/models/staticpagemodel"
	"shared/pkg/utils/apiutil"
	"shared/pkg/utils/dbutil"
	"shared/pkg/utils/grpcutil"
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

func (h *StaticPageHandler) Create(c *gin.Context) {
	var payload staticpagemodel.StaticPage
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, apiutil.Response{
			Code:    apiutil.CODE_ERROR,
			Message: "invalid payload",
		})
		return
	}

	var req staticpagepb.CreateRequest
	if err := dbutil.MapStruct(payload, &req); err != nil {
		c.JSON(http.StatusBadRequest, apiutil.Response{
			Code:    apiutil.CODE_ERROR,
			Message: "invalid payload",
		})
		return
	}

	resp, err := h.client.Create(context.Background(), &req)
	if err != nil {
		bootstrap.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Warn("StaticPageHandler >>> Create: failed to create static page")

		fields, ok := grpcutil.ParseValidationError(err)
		fmt.Println("***********", ok)
		fmt.Println(fields)

		c.JSON(http.StatusBadRequest, apiutil.Response{
			Code:    apiutil.CODE_ERROR,
			Message: "An error happened when creating new static page",
		})
		return
	}

	c.JSON(http.StatusOK, apiutil.Response{
		Code:    apiutil.CODE_SUCCESS,
		Message: "Success",
		Data:    resp.Id,
	})
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

	var staticPage staticpagemodel.StaticPage
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

func (h *StaticPageHandler) List(c *gin.Context) {
	resp, err := h.client.List(context.Background(), &staticpagepb.ListRequest{})
	if err != nil {
		bootstrap.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Warn("StaticPageHandler >>> List: failed to get static pages")

		c.JSON(http.StatusBadRequest, apiutil.Response{
			Code:    apiutil.CODE_ERROR,
			Message: "An error happened when getting data",
		})
		return
	}

	var staticPages []*staticpagemodel.StaticPage

	for _, staticPagePB := range resp.GetPages() {
		var staticPage staticpagemodel.StaticPage
		if err := dbutil.MapStruct(staticPagePB, &staticPage); err != nil {
			bootstrap.Logger.WithFields(logrus.Fields{
				"err":          err,
				"staticPagePB": staticPagePB,
			}).Warn("StaticPageHandler >>> List: failed to get map struct proto to model")
			continue
		}
		staticPages = append(staticPages, &staticPage)
	}

	c.JSON(http.StatusOK, apiutil.Response{
		Code:    apiutil.CODE_SUCCESS,
		Message: "Success",
		Data:    staticPages,
	})
}
