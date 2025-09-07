package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	v1 := r.Group("/v1") //version 1
	{
		v1.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "test API",
			})
		})
	}
}
