package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yungen-lu/shared-key-value-list-system/db/pgcodegen"
	"github.com/yungen-lu/shared-key-value-list-system/internal/domain"
)

//	type ListRepo interface {
//		GetByID(ctx context.Context, id int32) (List, error)
//		GetAll(ctx context.Context) ([]List, error)
//		Store(ctx context.Context, list List) error
//		DeleteByID(ctx context.Context, id int32) error
//	}
type ListRepo struct {
	queries *pgcodegen.Queries
}

func NewListRepo(pool *pgxpool.Pool) *ListRepo {
	// Pool         *pgxpool.Pool
	return &ListRepo{
		queries: pgcodegen.New(pool),
	}

}
func (l *ListRepo) GetByID(ctx context.Context, id int32) (domain.List, error) {
	var outputList domain.List
	list, err := l.queries.GetListByID(ctx, id)
	if err != nil {
		return outputList, err
	}
	outputList.ID = list.ID
	outputList.NextPageID = list.NextPageID.Int32
	outputList.PageCount = uint16(list.PageCount)
	return outputList, err
}
func (l *ListRepo) GetAll(ctx context.Context) ([]domain.List, error) {
	lists, err := l.queries.ListLists(ctx)
	if err != nil {
		return nil, err
	}
	outputLists := make([]domain.List, len(lists))
	for i := 0; i < len(lists); i++ {
		outputLists[i].ID = lists[i].ID
		outputLists[i].NextPageID = lists[i].NextPageID.Int32
		outputLists[i].PageCount = uint16(lists[i].PageCount)
	}
	return outputLists, err
}

func (l *ListRepo) Store(ctx context.Context, list domain.List) (int32, error) {
	r, err := l.queries.CreateList(ctx, pgcodegen.CreateListParams{})
	if err != nil {
		return 0, err
	}
	return r.ID, nil
}

func (l *ListRepo) DeleteByID(ctx context.Context, id int32) error {
	return nil
}
