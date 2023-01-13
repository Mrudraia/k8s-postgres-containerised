package models

import (
	"time"

	"gorm.io/gorm"
)

type Services struct {
	ID        uint      `gorm:"primary key;autoIncrement" json:"id"`
	Name      string    `gorm:"size:255;column:name;" json:"name"`
	Type      string    `gorm:"size:255;column:type;" json:"type"`
	Count     string    `gorm:"size:255;column:count;" json:"count"`
	CreatedAt time.Time `json:"created_at"`
}

func MigrateServices(db *gorm.DB) error {
	err := db.AutoMigrate(&Services{})
	return err
}
