package web

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	app "github.com/hsyntzgl/to-doList-Go/internal/app/user"
)

func HandlerError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, app.ErrInvalidCredentials):
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	case errors.Is(err, app.ErrEmailAllreadyExists):
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	case errors.Is(err, app.ErrForbidden):
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	case errors.Is(err, app.ErrUserNotFound):
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Sunucuda beklenmeyen bir hata oluştu"})
	}
}
