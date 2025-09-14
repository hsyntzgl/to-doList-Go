package main

import (
	app "github.com/hsyntzgl/to-doList-Go/internal/app/user"
	"github.com/hsyntzgl/to-doList-Go/internal/handlers"
	handler "github.com/hsyntzgl/to-doList-Go/internal/handlers/user"
	repository "github.com/hsyntzgl/to-doList-Go/internal/repository/user"
	"github.com/hsyntzgl/to-doList-Go/pkg/database"
	"github.com/hsyntzgl/to-doList-Go/pkg/utils/hasher"
	"github.com/hsyntzgl/to-doList-Go/pkg/utils/jwt"
)

func main() {
	jwtSecretKey := "secret_key"

	db := database.ConnectDB()

	userRepo := repository.NewPostgresUserRepository(db)
	hasher := hasher.NewBcryptHasher()
	tokenGenerator := jwt.NewTokenGenerator(jwtSecretKey)

	userService := app.NewUserService(userRepo, hasher, tokenGenerator)
	userHandler := handler.NewUserHandler(userService)

	router := handlers.SetupRoutes(userHandler, jwtSecretKey)
	router.Run()
}
