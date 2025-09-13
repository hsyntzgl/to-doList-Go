package user

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	app "github.com/hsyntzgl/to-doList-Go/internal/app/user"
)

type UserHandler struct {
	userService app.UserService
}

func NewUserHandler(us app.UserService) *UserHandler {
	return &UserHandler{userService: us}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req app.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gerçersiz istek: " + err.Error()})
		return
	}

	createdUser, err := h.userService.Register(c.Request.Context(), req.Username, req.Email, req.Password)

	if err != nil {
		if errors.Is(err, app.ErrEmailAllreadyExists) {
			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Sunucu Hatası", "message ": err.Error()})
		return
	}
	response := ToUserResponse(createdUser)
	c.JSON(http.StatusCreated, response)
}
func (h *UserHandler) Login(c *gin.Context) {
	var req app.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz istek: " + err.Error()})
		return
	}

	token, err := h.userService.Login(c.Request.Context(), req.Email, req.Password)

	if err != nil {
		if errors.Is(err, app.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Sunucu Hatası", "message ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, token)
}
func (h *UserHandler) UpdateUser(c *gin.Context) {
	var req app.UpdateUser
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz istek: " + err.Error()})
		return
	}

}
