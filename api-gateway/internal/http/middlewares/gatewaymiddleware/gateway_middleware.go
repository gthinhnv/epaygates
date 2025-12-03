package gatewaymiddleware

import (
	"apigateway/internal/bootstrap"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func ContextSetup() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Vary", "Origin, Access-Control-Request-Method, Access-Control-Request-Headers")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Lang, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE")
		c.Header("Access-Control-Max-Age", "86400")

		if c.Request.Method == http.MethodOptions {
			if origin := c.GetHeader("origin"); origin != "" {
				if isAllowedOrigin(origin) {
					c.Header("Access-Control-Allow-Origin", origin)
					c.Header("Access-Control-Allow-Credentials", "true")
					c.AbortWithStatus(http.StatusNoContent)
					return
				}
			}
		}

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

func isAllowedOrigin(origin string) bool {
	parsed, err := url.Parse(origin)
	if err != nil {
		return false
	}
	host := parsed.Hostname()
	allowedHosts := bootstrap.Config.AllowedHosts

	for _, h := range allowedHosts {
		if strings.HasPrefix(h, "*.") {
			// Wildcard match
			domain := strings.TrimPrefix(h, "*.")
			if host == domain || strings.HasSuffix(host, "."+domain) {
				return true
			}
		} else if h == host {
			// Exact match
			return true
		}
	}
	return false
}
