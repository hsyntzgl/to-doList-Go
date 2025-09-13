package entities

import "time"

type Tag struct {
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Tasks     []Task
}
