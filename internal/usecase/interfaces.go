package usecase

import (
	"context"

	"github.com/yungen-lu/shared-key-value-list-system/internal/domain"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test
type List interface {
	GetHeads(ctx context.Context) ([]domain.List, error)
	GetHeadByID(ctx context.Context, id int32) (domain.List, error)
	GetHeadByKey(ctx context.Context, key string) (domain.List, error)
	GetPages(ctx context.Context) ([]domain.Page, error)
	GetPageByID(ctx context.Context, id int32) (domain.Page, error)
	GetPageByKey(ctx context.Context, key string) (domain.Page, error)
	CreateHead(ctx context.Context, list domain.List) error
	CreatePage(ctx context.Context, page domain.Page) error
	UpdateHeadByKey(ctx context.Context, key string, list domain.List) error
	UpdatePageByKey(ctx context.Context, key string, page domain.Page) error
}
