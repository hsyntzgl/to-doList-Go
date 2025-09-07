package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hsyntzgl/to-doList-Go/pkg/database"
	"github.com/hsyntzgl/to-doList-Go/pkg/router"
)

func main() {
	r := gin.Default()

	database.ConnectDB()
	database.Migrate()

	router.SetupRoutes(r)

	if err := r.Run(":3000"); err != nil {
		log.Fatalf("web service cant start %v", err)
	}
}
