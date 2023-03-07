package domain

import (
	"context"
	"time"
)

type User struct {
	ID           int32
	Name         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	PersonalList List
}

type UserRepo interface {
	GetByID(ctx context.Context, id int32) (User, error)
	GetAll(ctx context.Context) ([]User, error)
	Store(ctx context.Context, user User) error
	DeleteByID(ctx context.Context, id int32) error
}
