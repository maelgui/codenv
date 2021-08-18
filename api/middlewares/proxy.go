package middlewares

import (
	"codenv-api/controllers"
	"codenv-api/utils"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func ReverseProxyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		proxy := regexp.MustCompile(`^(\d{4,5})-(\w{8}(?:-\w{4}){3}-\w{12})\.env.maelgui.fr$`)
		matches := proxy.FindStringSubmatch(c.Request.Host)

		if matches == nil {
			c.Next()
			return
		}

		remote, err := controllers.GetProxyRemote(matches[2], matches[1])
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "Proxy access forbidden"})
			c.Abort()
			return
		}

		utils.Proxy(c, remote, c.Request.URL.Path)

		c.Abort()
	}
}
