package database

import (
	"github.com/hsyntzgl/to-doList-Go/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	connStr := config.Config("connectionString")

	var err error

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  connStr,
		PreferSimpleProtocol: true,
	}), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})

	if err != nil {
		panic(err)
	}
}
