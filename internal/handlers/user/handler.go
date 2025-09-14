package user

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	app "github.com/hsyntzgl/to-doList-Go/internal/app/user"
	"github.com/hsyntzgl/to-doList-Go/internal/domain/repositories"
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
func (h *UserHandler) UpdateCurrentUser(c *gin.Context) {
	var req app.UpdateUser
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz istek: " + err.Error()})
		return
	}
	actorIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Yekilendirme Hatası"})
		return
	}

	actorID := actorIDInterface.(string)

	err := h.userService.UpdateUser(c.Request.Context(), actorID, actorID, req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Kullanıcı bilgileri güncellenirken hata meydana geldi", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kullanıcı bilgileri başarıyla güncellendi"})
}
func (h *UserHandler) DeleteCurrentUser(c *gin.Context) {
	actorIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Yetkilendirme Hatası"})
		return
	}

	actorID := actorIDInterface.(string)

	err := h.userService.Delete(c.Request.Context(), actorID, actorID)

	if err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Kullanıcı Bulunamadı"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Sunucu Hatası", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Hesap başarıyla silindi"})
}
