package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yungen-lu/shared-key-value-list-system/db/pgcodegen"
	"github.com/yungen-lu/shared-key-value-list-system/internal/domain"
)

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
	if page.NextID.Valid {
		outputPage.NextPageID = &page.NextID.Int32
	} else {
		outputPage.NextPageID = nil
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
		if pages[i].NextID.Valid {
			outputPages[i].NextPageID = &pages[i].NextID.Int32
		} else {
			outputPages[i].NextPageID = nil
		}
	}
	return outputPages, nil

}
func (p *PageRepo) Store(ctx context.Context, page domain.Page) (int32, error) {
	if page.NextPageID == nil {
		return 0, domain.ErrBadParamInput
	}
	r, err := p.queries.CreatePage(ctx, pgtype.Int4{Int32: *page.NextPageID, Valid: true})
	if err != nil {
		return 0, err
	}
	return r.ID, nil
}

func (p *PageRepo) DeleteByID(ctx context.Context, id int32) error {
	return nil
}
