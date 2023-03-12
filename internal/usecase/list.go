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

func NewListUseCase(listRepo domain.ListRepo, pageRepo domain.PageRepo) *ListUseCase {
	return &ListUseCase{
		l: listRepo,
		p: pageRepo,
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
