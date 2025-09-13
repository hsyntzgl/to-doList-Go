package user

import "github.com/hsyntzgl/to-doList-Go/internal/domain/entities"

func fromDomainUser(userEntity *entities.User) *User {
	return &User{
		ID:           userEntity.ID,
		Username:     userEntity.Username,
		Email:        userEntity.Email,
		PasswordHash: userEntity.PasswordHash,
		CreatedAt:    userEntity.CreatedAt,
		UpdatedAt:    userEntity.UpdatedAt,
	}
}
func toDomainUser(userModel *User) *entities.User {
	return &entities.User{
		ID:           userModel.ID,
		Username:     userModel.Username,
		Email:        userModel.Email,
		PasswordHash: userModel.PasswordHash,
		CreatedAt:    userModel.CreatedAt,
		UpdatedAt:    userModel.UpdatedAt,
	}
}
