package domain

import "context"

type List struct {
	ID          int32
	Key         string
	PageCount   uint16
	NextPageKey *string
}

//go:generate mockgen -source=list.go -destination=mocks/mock_list.go -package=mocks
type ListRepo interface {
	GetByID(ctx context.Context, id int32) (List, error)
	GetByKey(ctx context.Context, key string) (List, error)
	GetAll(ctx context.Context) ([]List, error)
	Store(ctx context.Context, list List) error
	UpdateByKey(ctx context.Context, key string, list List) error
	// Update(ctx context.Context, list List) error
	DeleteByID(ctx context.Context, id int32) error
}
