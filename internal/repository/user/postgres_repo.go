package user

import (
	"context"

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
	return nil
}
func (r *postgresUserRepository) Update(ctx context.Context, user *entities.User) error {
	return nil
}
func (r *postgresUserRepository) GetByID(ctx context.Context, id string) (*entities.User, error) {
	return nil, nil
}
func (r *postgresUserRepository) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	return nil, nil
}
func (r *postgresUserRepository) Delete(ctx context.Context, id string) error {
	return nil
}
