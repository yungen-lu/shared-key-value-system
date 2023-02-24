package domain

import "time"

type User struct {
	ID           uint64
	Name         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	PersonalList List
}
