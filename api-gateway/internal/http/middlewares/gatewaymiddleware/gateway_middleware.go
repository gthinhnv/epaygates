package gatewaymiddleware

import (
	"apigateway/internal/bootstrap"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type AllowedRule struct {
	Exact    map[string]struct{}
	Wildcard []string // suffixes like "example.com"
}

var allowedOriginCache sync.Map // map[string]bool
var allowedRule *AllowedRule
var allowedRuleOnce sync.Once

func ContextSetup() gin.HandlerFunc {
	allowedRuleOnce.Do(func() {
		allowedRule = initAllowedHosts(bootstrap.Config.AllowedHosts)
	})
	return func(c *gin.Context) {
		c.Header("Vary", "Origin, Access-Control-Request-Method, Access-Control-Request-Headers")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Lang, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE")
		c.Header("Access-Control-Max-Age", "86400")

		if origin := c.GetHeader("origin"); origin != "" {
			if isAllowedOrigin(origin) {
				c.Header("Access-Control-Allow-Origin", origin)
				c.Header("Access-Control-Allow-Credentials", "true")
			}
		}

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

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

func isAllowedOrigin(origin string) bool {
	if v, ok := allowedOriginCache.Load(origin); ok {
		return v.(bool)
	}

	// Faster than url.Parse; avoids allocations.
	// origin is typically: "https://example.com"
	host := extractHost(origin)
	if host == "" {
		allowedOriginCache.Store(origin, false)
		return false
	}

	// Exact match (O(1))
	if _, ok := allowedRule.Exact[host]; ok {
		allowedOriginCache.Store(origin, true)
		return true
	}

	// Wildcard matches
	for _, w := range allowedRule.Wildcard {
		// if host == w OR ends with .w
		if host == w || strings.HasSuffix(host, "."+w) {
			allowedOriginCache.Store(origin, true)
			return true
		}
	}

	allowedOriginCache.Store(origin, false)
	return false
}

func initAllowedHosts(hosts []string) *AllowedRule {
	rule := AllowedRule{
		Exact:    make(map[string]struct{}, len(hosts)),
		Wildcard: make([]string, 0),
	}

	for _, h := range hosts {
		if strings.HasPrefix(h, "*.") {
			rule.Wildcard = append(rule.Wildcard, h[2:])
		} else {
			rule.Exact[h] = struct{}{}
		}
	}

	return &rule
}

func extractHost(origin string) string {
	// Find "://"
	i := strings.Index(origin, "://")
	if i < 0 {
		return ""
	}
	rest := origin[i+3:]

	// Host ends at first '/' (if any)
	if slash := strings.IndexByte(rest, '/'); slash >= 0 {
		rest = rest[0:slash]
	}

	// Remove port, if exists
	if colon := strings.IndexByte(rest, ':'); colon >= 0 {
		return rest[:colon]
	}

	return rest
}
