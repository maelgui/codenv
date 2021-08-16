package schema

import "codenv-api/models"

type CreateWorkspace struct {
	Name  string `json:"name" binding:"required"`
	Image string `json:"image" binding:"required"`
}

type ResponseWorkspace struct {
	models.Workspace
	Status string `json:"status"`
}
