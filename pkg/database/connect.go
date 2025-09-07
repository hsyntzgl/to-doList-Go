package database

import (
	"database/sql"

	"github.com/hsyntzgl/to-doList-Go/pkg/config"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	connStr := config.Config("connectionString")

	var err error

	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}
}
