package user

import (
	"context"
	"errors"

	"github.com/hsyntzgl/to-doList-Go/internal/domain/entities"
	"github.com/hsyntzgl/to-doList-Go/internal/domain/repositories"
	"gorm.io/gorm"
)

type postgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) repositories.UserRepository {
	return &postgresUserRepository{
		db: db,
	}
}

func (r *postgresUserRepository) Create(ctx context.Context, user *entities.User) error {
	userModel := fromDomainUser(user)

	result := r.db.WithContext(ctx).Create(&userModel)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
func (r *postgresUserRepository) Update(ctx context.Context, user *entities.User) error {
	userModel := fromDomainUser(user)

	result := r.db.WithContext(ctx).Where("id = ?", userModel.ID).Updates(userModel)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return repositories.ErrNotFound
	}

	return nil
}
func (r *postgresUserRepository) GetByID(ctx context.Context, id string) (*entities.User, error) {

	var userModel User

	result := r.db.WithContext(ctx).Where("id = ?", id).First(&userModel)

	if result.Error != nil {
		if errors.Is(result.Error, repositories.ErrNotFound) {
			return nil, repositories.ErrNotFound
		}
		return nil, result.Error
	}

	return toDomainUser(&userModel), nil
}
func (r *postgresUserRepository) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	var userModel User

	result := r.db.WithContext(ctx).Where("email = ?", email).First(&userModel)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, repositories.ErrNotFound
		}
		return nil, result.Error
	}

	return toDomainUser(&userModel), nil
}
func (r *postgresUserRepository) Delete(ctx context.Context, id string) error {

	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&User{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return repositories.ErrNotFound
	}

	return nil
}
