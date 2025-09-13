package entities

import "time"

type Task struct {
	ID          uint
	Title       string
	Description string
	IsCompleted bool
	Priority    int
	DueDate     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserID      string
	User        string
}
