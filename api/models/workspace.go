package models

type Workspace struct {
	*BaseModel
	ContainerID string `json:"container_id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
}

const (
	StatusPending = "PENDING"
	StatusRunning = "RUNNING"
	StatusStopped = "STOPPED"
)
