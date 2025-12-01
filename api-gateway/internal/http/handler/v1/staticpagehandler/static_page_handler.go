package staticpagehandler

import (
	"apigateway/gen/go/staticpagepb"
	"apigateway/internal/bootstrap"
	"context"
	"net/http"

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
	resp, err := h.client.Get(context.Background(), &staticpagepb.GetRequest{
		Id: 1,
	})
	if err != nil {
		bootstrap.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Warn("StaticPageHandler >>> Get: failed to get static page")

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get static page",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    resp.Page,
	})
}
