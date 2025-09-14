package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hsyntzgl/to-doList-Go/internal/handlers/user"
	"github.com/hsyntzgl/to-doList-Go/internal/middleware"
)

func SetupRoutes(userhandler *user.UserHandler, jwtSecretKey string) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		//public user endpoints
		usersPublic := v1.Group("/users")
		{
			usersPublic.POST("/register", userhandler.Register)
			usersPublic.POST("/login", userhandler.Login)
		}

		authRequired := v1.Group("/")
		authRequired.Use(middleware.AuthMiddleware(jwtSecretKey))
		{
			users := authRequired.Group("/users")
			{
				users.POST("/update-user", userhandler.UpdateCurrentUser)
				users.DELETE("/delete-account", userhandler.DeleteCurrentUser)
			}
		}
	}
	return router
}
