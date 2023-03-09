package usecase

import (
	"context"

	"github.com/yungen-lu/shared-key-value-list-system/internal/domain"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test
type List interface {
	GetHeads(ctx context.Context) ([]domain.List, error)
	GetHeadByID(ctx context.Context, id int32) (domain.List, error)
	GetPages(ctx context.Context) ([]domain.Page, error)
	GetPageByID(ctx context.Context, id int32) (domain.Page, error)
	// GetHead(id int32) (domain.Page, error)
	// GetPage(id int32) (domain.Page, error)
}
