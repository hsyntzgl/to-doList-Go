package user

import (
	app "github.com/hsyntzgl/to-doList-Go/internal/app/user"
	"github.com/hsyntzgl/to-doList-Go/internal/domain/entities"
)

func ToUserResponse(user *entities.User) app.UserResponse {
	return app.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}
