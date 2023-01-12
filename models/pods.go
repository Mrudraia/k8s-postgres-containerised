package models

import (
	"time"

	"gorm.io/gorm"
)

type Pods struct {
	ID        uint      `gorm:"primary key;autoIncrement" json:"id"`
	Name      string    `json:"name"`
	Namespace string    `json:"namespace"`
	Count     string    `json:"count"`
	CreatedAt time.Time `json:"created_at"`
}

func MigratePods(db *gorm.DB) error {
	err := db.AutoMigrate(&Pods{})
	return err
}
