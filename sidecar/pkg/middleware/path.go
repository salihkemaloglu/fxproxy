package middleware

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

// ValidatePathMiddleware checks paths is allowed
func ValidatePathMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !ValidatePath(c.Request.URL.String()) {
			c.AbortWithStatus(http.StatusNotAcceptable)
			return
		}

		c.Next()
	}
}

// ValidatePath validates paths
func ValidatePath(path string) bool {
	for _, al := range getAllowedList() {
		al = strings.TrimPrefix(al, "/")
		al = strings.TrimSuffix(al, "/")

		path = strings.TrimPrefix(path, "/")
		if path == al {
			return true
		}

		if regexp.MustCompile(generateRegexQuery(al)).MatchString(path) {
			return true
		}
	}

	return false
}

// getAllowedList return allowed paths
func getAllowedList() []string {
	return []string{
		"/hello/",
		"/company/",
		"/company/{id}",
		"/company/account",
		"/account",
		"/account/{id}",
		"/{id}",
		"/account/{id}/user",
		"/tenant/account/blocked",
	}
}

// generateRegexQuery generates regex query to check path
func generateRegexQuery(allowedPath string) string {
	alpParts := strings.Split(allowedPath, "/")
	reQuote := "^"
	for i := range alpParts {
		if alpParts[i] == "{id}" {
			reQuote += "(sd|sj|acc)[a-z0-9]+/"
		} else {
			reQuote += alpParts[i] + "/"
		}
	}

	return reQuote[:len(reQuote)-1] + "$"
}
