package usecase

import (
	"context"
	"time"

	"github.com/yungen-lu/shared-key-value-list-system/internal/domain"
)

type ListUseCase struct {
	l              domain.ListRepo
	p              domain.PageRepo
	contextTimeout time.Duration
}

var _ List = (*ListUseCase)(nil)

func NewListUseCase(listRepo domain.ListRepo, pageRepo domain.PageRepo, timeout time.Duration) *ListUseCase {
	return &ListUseCase{
		l:              listRepo,
		p:              pageRepo,
		contextTimeout: timeout,
	}
}

func (uc *ListUseCase) GetHeads(ctx context.Context) ([]domain.List, error) {
	c, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()
	return uc.l.GetAll(c)
}

func (uc *ListUseCase) GetHeadByID(ctx context.Context, id int32) (domain.List, error) {
	c, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()
	return uc.l.GetByID(c, id)
}

func (uc *ListUseCase) GetHeadByKey(ctx context.Context, key string) (domain.List, error) {
	c, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()
	return uc.l.GetByKey(c, key)
}

func (uc *ListUseCase) GetPages(ctx context.Context) ([]domain.Page, error) {
	c, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()
	return uc.p.GetAll(c)
}

func (uc *ListUseCase) GetPageByID(ctx context.Context, id int32) (domain.Page, error) {
	c, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()
	return uc.p.GetByID(c, id)
}

func (uc *ListUseCase) GetPageByKey(ctx context.Context, key string) (domain.Page, error) {
	c, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()
	return uc.p.GetByKey(c, key)
}

func (uc *ListUseCase) CreateHead(ctx context.Context, list domain.List) error {
	c, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()
	return uc.l.Store(c, list)
}

func (uc *ListUseCase) CreatePage(ctx context.Context, page domain.Page) error {
	c, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()
	return uc.p.Store(c, page)
}

// CreatePage(ctx context.Context, page domain.Page) error

func (uc *ListUseCase) UpdateHeadByKey(ctx context.Context, key string, list domain.List) error {
	c, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()
	return uc.l.UpdateByKey(c, key, list)
}

func (uc *ListUseCase) UpdatePageByKey(ctx context.Context, key string, page domain.Page) error {
	c, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()
	return uc.p.UpdateByKey(c, key, page)
}
