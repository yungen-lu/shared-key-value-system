package domain

import (
	"context"
	"time"
)

type Topic struct {
	ID          int32
	Name        string
	Description string
	PopularList List
	NewestList  List
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

//go:generate mockgen -source=topic.go -destination=mocks/mock_topic.go -package=mocks
type TopicRepo interface {
	GetByID(id int32) (Topic, error)
	GetAll(ctx context.Context) ([]Topic, error)
	Store(ctx context.Context, topic Topic) error
	DeleteByID(ctx context.Context, id int32) error
}
