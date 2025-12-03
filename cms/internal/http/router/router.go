package router

import (
	"cms/internal/http/handler/dashboardhandler"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	if os.Getenv("APP_ENV") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	dashboardHandler := dashboardhandler.NewStaticPageHandler()

	r := gin.New()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok1"})
	})

	r.GET("/dashboard", dashboardHandler.GetIndex)

	return r
}
