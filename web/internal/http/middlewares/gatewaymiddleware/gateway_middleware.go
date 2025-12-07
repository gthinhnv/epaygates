package gatewaymiddleware

import (
	"strconv"
	"time"
	"web/internal/bootstrap"

	"github.com/gin-gonic/gin"
)

func ContextSetup() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("lang", bootstrap.Config.DefaultLang)

		c.Next()
	}
}

func StaticCache(lifeTime time.Duration) gin.HandlerFunc {
	lifeTimeSt := strconv.Itoa(int(lifeTime.Seconds()))
	return func(c *gin.Context) {
		c.Header("Cache-Control", "max-age="+lifeTimeSt)
		c.Next()
	}
}
