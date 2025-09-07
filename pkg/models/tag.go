package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name  string `gorm:"unique;not null"`
	Todos []Todo `gorm:"many2many:todo_tags"`
}
