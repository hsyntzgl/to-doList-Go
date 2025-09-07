package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	IsCompleted bool `gorm:"default:false"`
	Priority    int  `gorm:"default:0"`
	DueDate     *time.Time
	UserID      uint
	User        User
}
