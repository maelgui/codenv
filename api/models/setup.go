package models

import (
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open(mysql.Open(os.Getenv("DATABASE_CONF")), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Workspace{})

	DB = database
}
