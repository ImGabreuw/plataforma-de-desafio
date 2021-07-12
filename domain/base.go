package domain

import "time"

type Base struct {
	ID        string
	CreatedAt time.Timer
	UpdatedAt time.Timer
	DeletedAt time.Timer
}
