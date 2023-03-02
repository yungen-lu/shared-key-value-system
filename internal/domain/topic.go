package domain

import "time"

type Topic struct {
	ID          int32
	Name        string
	Description string
	PopularList List
	NewestList  List
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TopicRepo interface {
	GetByID(id int32) (Topic, error)
}
