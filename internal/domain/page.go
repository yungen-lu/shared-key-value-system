package domain

import "context"

type Page struct {
	ID         int32 // key
	Articles   []Article
	NextPageID int32
}

//go:generate mockgen -source=page.go -destination=mocks/mock_page.go -package=mocks
type PageRepo interface {
	GetByID(ctx context.Context, id int32) (Page, error)
	GetAll(ctx context.Context) ([]Page, error)
	Store(ctx context.Context, page Page) error
	DeleteByID(ctx context.Context, id int32) error
}
