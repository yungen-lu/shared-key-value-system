package domain

import "time"

type User struct {
	ID           int32
	Name         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	PersonalList List
}

type UserRepo interface {
	GetByID(id int32) (User, error)
}
