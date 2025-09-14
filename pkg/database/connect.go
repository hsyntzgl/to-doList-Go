package database

import (
	"log"

	"github.com/hsyntzgl/to-doList-Go/internal/repository/user"
	"github.com/hsyntzgl/to-doList-Go/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	connStr := config.Config("connectionString")

	var err error

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connStr,
		PreferSimpleProtocol: true,
	}), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})

	if err := db.AutoMigrate(&user.User{}); err != nil {
		log.Fatal(err.Error())
	}

	if err != nil {
		panic(err)
	}

	return db
}
