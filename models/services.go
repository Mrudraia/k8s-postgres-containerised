package models

import (
	"time"

	"gorm.io/gorm"
)

type Services struct {
	ID        uint      `gorm:"primary key;autoIncrement" json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Count     string    `json:"count"`
	CreatedAt time.Time `json:"created_at"`
}

func MigrateServices(db *gorm.DB) error {
	err := db.AutoMigrate(&Services{})
	return err
}
