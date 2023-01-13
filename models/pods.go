package models

import (
	"time"

	"gorm.io/gorm"
)

type Pods struct {
	ID        uint      `gorm:"primary key;autoIncrement" json:"id"`
	Name      string    `gorm:"size:255;column:name;" json:"name"`
	Namespace string    `gorm:"size:255;column:namespace;" json:"namespace"`
	Count     string    `gorm:"size:255;column:count;" json:"count"`
	CreatedAt time.Time `json:"created_at"`
}

func MigratePods(db *gorm.DB) error {
	err := db.AutoMigrate(&Pods{})
	return err
}
