package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yungen-lu/shared-key-value-list-system/db/pgcodegen"
	"github.com/yungen-lu/shared-key-value-list-system/internal/domain"
)

var _ domain.ListRepo = (*ListRepo)(nil)

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
	outputList.Key = list.Key
	outputList.PageCount = uint16(list.PageCount)
	if list.NextPageKey.Valid {
		outputList.NextPageKey = &list.NextPageKey.String
	} else {
		outputList.NextPageKey = nil
	}
	return outputList, err
}

func (l *ListRepo) GetByKey(ctx context.Context, key string) (domain.List, error) {
	var outputList domain.List
	list, err := l.queries.GetListByKey(ctx, key)
	if err != nil {
		return outputList, err
	}
	outputList.ID = list.ID
	outputList.Key = list.Key
	outputList.PageCount = uint16(list.PageCount)
	if list.NextPageKey.Valid {
		outputList.NextPageKey = &list.NextPageKey.String
	} else {
		outputList.NextPageKey = nil
	}
	return outputList, nil
}

func (l *ListRepo) GetAll(ctx context.Context) ([]domain.List, error) {
	lists, err := l.queries.ListLists(ctx)
	if err != nil {
		return nil, err
	}
	outputLists := make([]domain.List, len(lists))
	for i := 0; i < len(lists); i++ {
		outputLists[i].ID = lists[i].ID
		outputLists[i].Key = lists[i].Key
		if lists[i].NextPageKey.Valid {
			outputLists[i].NextPageKey = &lists[i].NextPageKey.String
		} else {
			outputLists[i].NextPageKey = nil
		}
		outputLists[i].PageCount = uint16(lists[i].PageCount)
	}
	return outputLists, err
}

func (l *ListRepo) Store(ctx context.Context, list domain.List) error {
	var param pgcodegen.CreateListParams
	param.Key = list.Key
	if list.NextPageKey != nil {
		param.NextPageKey = pgtype.Text{String: *list.NextPageKey, Valid: true}
	} else {
		param.NextPageKey = pgtype.Text{Valid: false}
	}
	_, err := l.queries.CreateList(ctx, param)
	return err
}

func (l *ListRepo) DeleteByID(ctx context.Context, id int32) error {
	return nil
}

func (l *ListRepo) DeleteByKey(ctx context.Context, key string) error {
	count, err := l.queries.DeleteListByKey(ctx, key)
	if err != nil {
		return err
	}
	if count == 0 {
		return domain.ErrNotFound
	}
	return nil
}

func (l *ListRepo) UpdateByKey(ctx context.Context, key string, list domain.List) error {
	var param pgcodegen.UpdateListByKeyParams
	param.Key = key
	if list.NextPageKey != nil {
		param.NextPageKey = pgtype.Text{String: *list.NextPageKey, Valid: true}
	} else {
		param.NextPageKey = pgtype.Text{Valid: false}
	}
	_, err := l.queries.UpdateListByKey(ctx, param)
	return err
}
