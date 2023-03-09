package domain

import "context"

type List struct {
	ID         int32
	PageCount  uint16
	NextPageID int32
}

//go:generate mockgen -source=list.go -destination=mocks/mock_list.go -package=mocks
type ListRepo interface {
	GetByID(ctx context.Context, id int32) (List, error)
	GetAll(ctx context.Context) ([]List, error)
	Store(ctx context.Context, list List) error
	DeleteByID(ctx context.Context, id int32) error
}
