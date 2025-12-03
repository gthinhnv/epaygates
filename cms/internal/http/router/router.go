package router

import (
	"cms/internal/http/handlers/dashboardhandler"
	"cms/internal/http/middlewares/authmiddleware"
	"cms/internal/http/middlewares/gatewaymiddleware"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	if os.Getenv("APP_ENV") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	dashboardHandler := dashboardhandler.NewStaticPageHandler()

	r := gin.New()

	staticRouteGroup := r.Group("/", gatewaymiddleware.StaticCache(356*24*time.Hour))

	/*
	 * Serve static files
	 */
	staticRouteGroup.Static("/assets", "./assets")

	r.Use(gatewaymiddleware.ContextSetup())
	r.Use(authmiddleware.Authenticate())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok1"})
	})

	r.GET("/dashboard", dashboardHandler.GetIndex)

	return r
}
