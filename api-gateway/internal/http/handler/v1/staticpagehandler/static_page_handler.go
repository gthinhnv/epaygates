package staticpagehandler

import (
	"apigateway/gen/go/staticpagepb"
	"apigateway/internal/bootstrap"
	"context"
	"fmt"
	"net/http"
	"shared/models/staticpage"

	"github.com/mitchellh/mapstructure"

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
	fmt.Println("**************")
	staticPageDB := staticpage.StaticPage{
		Id:      1,
		Title:   "About Us",
		Content: "This is the About Us page content.",
		Status:  1,
	}
	var staticPageProto staticpagepb.StaticPage

	err := mapstructure.Decode(&staticPageDB, &staticPageProto)
	if err != nil {
		fmt.Println("Error decoding:", err)
		// return
	}
	fmt.Println("staticPageProto:", staticPageProto)
	return

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
