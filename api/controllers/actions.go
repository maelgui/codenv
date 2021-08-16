package controllers

import (
	"codenv-api/docker"
	"codenv-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartContainer(c *gin.Context) {
	var workspace models.Workspace

	if err := models.DB.Where("id = ?", c.Param("id")).First(&workspace).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	go docker.StartContainer(workspace.ContainerID)
}

func RestartContainer(c *gin.Context) {
	var workspace models.Workspace

	if err := models.DB.Where("id = ?", c.Param("id")).First(&workspace).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	go docker.RestartContainer(workspace.ContainerID)
}

func StopContainer(c *gin.Context) {
	var workspace models.Workspace

	if err := models.DB.Where("id = ?", c.Param("id")).First(&workspace).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	go docker.StopContainer(workspace.ContainerID)
}
