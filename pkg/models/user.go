package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Username string    `gorm:"unique;not null"`
	Email    string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
	Todos    []Todo
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}
