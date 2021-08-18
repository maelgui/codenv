package controllers

import (
	"codenv-api/docker"
	"codenv-api/models"
	"codenv-api/schema"
	"codenv-api/utils"
	"errors"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

func CreateProxy(c *gin.Context) {
	var input schema.CreateProxy
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create workspace
	proxy := models.Proxy{
		Name:        input.Name,
		Port:        input.Port,
		WorkspaceID: input.WorkspaceID,
	}
	models.DB.Create(&proxy)

	c.JSON(http.StatusOK, proxy)
}

func Proxy(c *gin.Context) {

	remote, err := GetProxyRemote(c.Param("id"), c.Param("id"))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Proxy access forbidden"})
	}

	utils.Proxy(c, remote)
}

func GetProxyRemote(workspaceID string, port string) (*url.URL, error) {
	var proxy models.Proxy
	if err := models.DB.Where("port = ?", port).Where("workspace_id = ?", workspaceID).First(&proxy).Error; err != nil {
		return nil, errors.New("Record not found!")
	}

	containerInfo := docker.RetrieveContainer(proxy.Workspace.ContainerID)
	ip := containerInfo.NetworkSettings.Networks[os.Getenv("DOCKER_NETWORK")].IPAddress

	target_url := "http://" + ip + ":" + port

	return url.Parse(target_url)
}
