package domain

import "time"

type Topic struct {
	ID          uint64
	Name        string
	Description string
	PopularList List
	NewestList  List
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TopicRepo interface {
	GetByID(id uint64) (Topic, error)
}
