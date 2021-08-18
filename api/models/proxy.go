package models

type Proxy struct {
	BaseModel
	Name        string    `json:"name"`
	Port        uint      `json:"port"`
	WorkspaceID string    `json:"workspace_id"`
	Workspace   Workspace `json:"-"`
}
