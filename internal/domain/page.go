package domain

import "context"

type Page struct {
	ID          int32
	Key         string
	Articles    []Article
	NextPageKey *string
}

//go:generate mockgen -source=page.go -destination=mocks/mock_page.go -package=mocks
type PageRepo interface {
	GetByID(ctx context.Context, id int32) (Page, error)
	GetByKey(ctx context.Context, key string) (Page, error)
	GetAll(ctx context.Context) ([]Page, error)
	Store(ctx context.Context, page Page) error
	UpdateByKey(ctx context.Context, key string, page Page) error
	DeleteByID(ctx context.Context, id int32) error
}
