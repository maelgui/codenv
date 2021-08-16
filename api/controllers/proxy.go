package controllers

import (
	"codenv-api/docker"
	"codenv-api/models"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func Proxy(c *gin.Context) {
	var workspace models.Workspace

	if err := models.DB.Where("id = ?", c.Param("id")).First(&workspace).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	containerInfo := docker.RetrieveContainer(workspace.ContainerID)
	ip := containerInfo.NetworkSettings.IPAddress

	// if err := models.DB.Where("id = ?", c.Param("id")).First(&workspace).Error; err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
	// 	return
	// }

	port := c.Param("port")

	remote, err := url.Parse("http://" + ip + ":" + port)
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	//Define the director func
	//This is a good place to log, for example
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("path")
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}
