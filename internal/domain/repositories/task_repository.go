package repositories

import (
	"context"

	"github.com/hsyntzgl/to-doList-Go/internal/domain/entities"
)

type TaskRepository interface {
	Create(ctx context.Context, task *entities.Task) error

	Update(ctx context.Context, task *entities.Task) error

	GetUserTasks(ctx context.Context, userID string) (*[]entities.Task, error)

	GetDetail(ctx context.Context, ID uint) (*entities.Task, error)

	GetAll(ctx context.Context) (*[]entities.Task, error)

	Delete(ctx context.Context, ID uint) error
}
