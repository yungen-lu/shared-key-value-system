package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yungen-lu/shared-key-value-list-system/db/pgcodegen"
	"github.com/yungen-lu/shared-key-value-list-system/internal/domain"
)

var _ domain.PageRepo = (*PageRepo)(nil)

type PageRepo struct {
	queries *pgcodegen.Queries
}

func NewPageRepo(pool *pgxpool.Pool) *PageRepo {
	return &PageRepo{
		queries: pgcodegen.New(pool),
	}
}

func (p *PageRepo) GetByID(ctx context.Context, id int32) (domain.Page, error) {
	var outputPage domain.Page
	page, err := p.queries.GetPageByID(ctx, id)
	if err != nil {
		return outputPage, err
	}
	outputPage.ID = page.ID
	// outputPage.Articles = page.Articles
	if page.NextPageKey.Valid {
		outputPage.NextPageKey = &page.NextPageKey.String
	} else {
		outputPage.NextPageKey = nil
	}
	return outputPage, nil
}

func (p *PageRepo) GetByKey(ctx context.Context, key string) (domain.Page, error) {
	var outputPage domain.Page
	page, err := p.queries.GetPageByKey(ctx, key)
	if err != nil {
		return outputPage, err
	}
	outputPage.ID = page.ID
	outputPage.Key = page.Key
	if page.NextPageKey.Valid {
		outputPage.NextPageKey = &page.NextPageKey.String
	} else {
		outputPage.NextPageKey = nil
	}
	return outputPage, nil
}

func (p *PageRepo) GetAll(ctx context.Context) ([]domain.Page, error) {
	pages, err := p.queries.ListPages(ctx)
	if err != nil {
		return nil, err
	}
	outputPages := make([]domain.Page, len(pages))
	for i := 0; i < len(pages); i++ {
		outputPages[i].ID = pages[i].ID
		// outputPages[i].Articles = pages[i].Articles
		if pages[i].NextPageKey.Valid {
			outputPages[i].NextPageKey = &pages[i].NextPageKey.String
		} else {
			outputPages[i].NextPageKey = nil
		}
	}
	return outputPages, nil

}
func (p *PageRepo) Store(ctx context.Context, page domain.Page) error {
	var param pgcodegen.CreatePageParams
	param.Key = page.Key
	if page.NextPageKey != nil {
		param.NextPageKey = pgtype.Text{String: *page.NextPageKey, Valid: true}
	} else {
		param.NextPageKey = pgtype.Text{Valid: false}
	}
	_, err := p.queries.CreatePage(ctx, param)
	return err
}

func (p *PageRepo) DeleteByID(ctx context.Context, id int32) error {
	return nil
}

func (p *PageRepo) UpdateByKey(ctx context.Context, key string, page domain.Page) error {
	var param pgcodegen.UpdatePageByKeyParams
	param.Key = key
	if page.NextPageKey != nil {
		param.NextPageKey = pgtype.Text{String: *page.NextPageKey, Valid: true}
	} else {
		param.NextPageKey = pgtype.Text{Valid: false}
	}
	_, err := p.queries.UpdatePageByKey(ctx, param)
	return err
}
