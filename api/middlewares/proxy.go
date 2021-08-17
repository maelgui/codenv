package middlewares

import (
	"codenv-api/docker"
	"codenv-api/models"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strconv"

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

		port, err := strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}

		ReverseProxy(c, matches[2], port)

		c.Abort()
	}
}

func ReverseProxy(c *gin.Context, workspaceID string, port int) {
	var workspace models.Workspace
	models.DB.First(&workspace, "id = ?", workspaceID)

	containerInfo := docker.RetrieveContainer(workspace.ContainerID)
	ip := containerInfo.NetworkSettings.IPAddress

	remote, err := url.Parse(fmt.Sprintf("http://%s:%d", ip, port))
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)

	proxy.Director = func(r *http.Request) {
		r.Header = c.Request.Header
		r.Host = remote.Host
		r.URL.Scheme = remote.Scheme
		r.URL.Host = remote.Host
		r.URL.Path = c.Request.URL.Path
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}
