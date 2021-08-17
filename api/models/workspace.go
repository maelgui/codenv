package models

type Workspace struct {
	*BaseModel
	ContainerID string `json:"container_id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
}

func (e *Workspace) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.NewString()
	return
}
