package authmiddleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Vary", "Origin, Access-Control-Request-Method, Access-Control-Request-Headers") // Required for proxies/caching
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Lang, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE")
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Max-Age", "86400") // 24h cache for preflights

		// Handle preflight
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
