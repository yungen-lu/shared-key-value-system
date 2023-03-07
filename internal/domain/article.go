package domain

import (
	"context"
	"time"
)

type Article struct {
	ID        int32
	Title     string
	Content   string
	Author    User
	UpdatedAt time.Time
	CreatedAt time.Time
	Topic     Topic
	Tags      string // TODO
}

type ArticleRepo interface {
	GetByID(ctx context.Context, id int32) (Article, error)
	GetByTitle(ctx context.Context, title string) (Article, error)
	// Update(ctx context.Context,article Article) error
	Store(ctx context.Context, article Article) error
	DeleteByID(ctx context.Context, id int32) error
}
