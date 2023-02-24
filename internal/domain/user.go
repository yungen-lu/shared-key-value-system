package domain

import "time"

type User struct {
	ID           uint64
	Name         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	PersonalList List
}

type UserRepo interface {
	GetByID(id uint64) (User, error)
}
