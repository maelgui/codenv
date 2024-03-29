package models

import (
	"os"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (e *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.NewString()
	return
}

var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_DNS")), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Workspace{}, &Proxy{})

	DB = database
}
