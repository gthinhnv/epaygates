package router

import (
	"apigateway/internal/http/router/routerv1"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	if os.Getenv("APP_ENV") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	/*
	 * V1 routes
	 */
	routerv1.New(r)

	return r
}
