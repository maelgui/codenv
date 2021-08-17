package schema

import "codenv-api/models"

type CreateWorkspace struct {
	Name  string `json:"name" binding:"required"`
	Image string `json:"image" binding:"required"`
}

type CreateProxy struct {
	Name        string `json:"name" binding:"required"`
	Port        uint   `json:"port" binding:"required"`
	WorkspaceID string `json:"workspace_id" binding:"required"`
}

type ResponseWorkspace struct {
	models.Workspace
	Status string `json:"status"`
}
