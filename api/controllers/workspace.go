package controllers

import (
	"codenv-api/docker"
	"codenv-api/models"
	"codenv-api/schema"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListWorkspace(c *gin.Context) {
	var workspaces []models.Workspace

	models.DB.Preload("Proxies").Find(&workspaces)

	var responses []schema.ResponseWorkspace

	for _, w := range workspaces {
		response := schema.ResponseWorkspace{Workspace: w}
		if w.ContainerID != "" {
			response.Status = docker.RetrieveContainer(w.ContainerID).State.Status
		}
		responses = append(responses, response)
	}

	c.JSON(http.StatusOK, responses)
}

func RetrieveWorkspace(c *gin.Context) {
	var workspace models.Workspace

	if err := models.DB.Where("id = ?", c.Param("id")).First(&workspace).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, workspace)
}

func CreateWorkspace(c *gin.Context) {
	var input schema.CreateWorkspace
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create workspace
	workspace := models.Workspace{
		Name:  input.Name,
		Image: input.Image,
	}

	models.DB.Create(&workspace)

	c.JSON(http.StatusOK, workspace)
}

func DeleteWorkspace(c *gin.Context) {
	var workspace models.Workspace

	if err := models.DB.Where("id = ?", c.Param("id")).First(&workspace).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&workspace)

	c.Status(http.StatusOK)

	go docker.DeleteContainer(workspace.ContainerID)
}
