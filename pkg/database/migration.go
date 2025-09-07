package database

import (
	"log"

	"github.com/hsyntzgl/to-doList-Go/pkg/models"
)

func Migrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Todo{},
		&models.Tag{},
	)

	if err != nil {
		log.Fatalf("Database migration failed %v", err)
	}
}
