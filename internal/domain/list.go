package domain

import "context"

type List struct {
	ID         int32
	PageCount  uint16
	NextPageID int32
}

type ListRepo interface {
	GetByID(ctx context.Context, id int32) (List, error)
	GetAll(ctx context.Context) ([]List, error)
	Store(ctx context.Context, list List) error
	DeleteByID(ctx context.Context, id int32) error
}
