package usecase

import (
	"context"

	"github.com/yungen-lu/shared-key-value-list-system/internal/domain"
)

type listUseCase struct {
	l domain.ListRepo
	p domain.PageRepo
}

var _ List = (*listUseCase)(nil)

func NewListUseCase() *listUseCase {
	return &listUseCase{}
}

func (uc *listUseCase) GetHeads(ctx context.Context) ([]domain.List, error) {
	return nil, nil
}

func (uc *listUseCase) GetHeadByID(ctx context.Context, id int32) (domain.List, error) {
	return domain.List{}, nil
}

func (uc *listUseCase) GetPages(ctx context.Context) ([]domain.Page, error) {
	return nil, nil
}

func (uc *listUseCase) GetPageByID(ctx context.Context, id int32) (domain.Page, error) {
	return domain.Page{}, nil
}
