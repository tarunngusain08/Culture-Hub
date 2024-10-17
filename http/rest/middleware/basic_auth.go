package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/tarunngusain08/culturehub/http/rest/handlers"
	"github.com/tarunngusain08/culturehub/http/rest/session"
	"github.com/tarunngusain08/culturehub/pkg/models"
)

func BasicAuthMiddleware(dao models.DaoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		prefix := []string{handlers.LoginPath, "/", "/hello/world"}
		contains := []string{}
		suffix := []string{".jpeg", ".jpg", ".css"}
		if anyTrue(c.Request.URL.Path, prefix, contains, suffix) {
			logger.Info("skipping basic auth")
			return
		}
		usr, pss, err := session.GetUserPass(c)
		if err != nil || usr == "" || pss == "" {
			c.Redirect(http.StatusTemporaryRedirect, handlers.LoginPath)
		}
	}
}

func anyTrue(s string, prefixs, contains, suffixes []string) bool {
	startsWith := func(s string, prefix string) bool {
		return len(s) >= len(prefix) && s[:len(prefix)] == prefix
	}

	endsWith := func(s string, suffix string) bool {
		return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
	}
	for _, suffix := range suffixes {
		if endsWith(s, suffix) {
			return true
		}
	}

	for _, contains := range contains {
		if strings.Contains(s, contains) {
			return true
		}
	}

	for _, prefix := range prefixs {
		if startsWith(s, prefix) {
			return true
		}
	}
	return false
}
