package usecase_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/yungen-lu/shared-key-value-list-system/internal/domain"
	"github.com/yungen-lu/shared-key-value-list-system/internal/domain/mocks"
	"github.com/yungen-lu/shared-key-value-list-system/internal/usecase"
)

// GetHeads(ctx context.Context) ([]domain.List, error)
// GetHeadByID(ctx context.Context, id int32) (domain.List, error)
// GetPages(ctx context.Context) ([]domain.Page, error)
// GetPageByID(ctx context.Context, id int32) (domain.Page, error)
type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

func list(t *testing.T) (*usecase.ListUseCase, *mocks.MockListRepo, *mocks.MockPageRepo) {
	t.Helper()
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	listRepo := mocks.NewMockListRepo(mockCtl)
	pageRepo := mocks.NewMockPageRepo(mockCtl)
	list := usecase.NewListUseCase(listRepo, pageRepo)
	return list, listRepo, pageRepo
}
func TestGetHeads(t *testing.T) {
	t.Parallel()

	listUseCase, listRepo, _ := list(t)
	tests := []test{
		{
			name: "error result",
			mock: func() {
				listRepo.EXPECT().GetAll(gomock.Any()).Return(nil, domain.ErrInernalServerError)
			},
			res: []domain.List(nil),
			err: domain.ErrInernalServerError,
		},
		{
			name: "success result",
			mock: func() {
				listRepo.EXPECT().GetAll(gomock.Any()).Return([]domain.List{}, nil)
			},
			res: []domain.List{},
			err: nil,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.mock()
			res, err := listUseCase.GetHeads(context.Background())
			assert.Equal(t, tc.res, res)
			assert.ErrorIs(t, tc.err, err)
		})
	}

}

func TestGetHeadByID(t *testing.T) {
	t.Parallel()

	listUseCase, listRepo, _ := list(t)
	tests := []test{
		{
			name: "error result",
			mock: func() {
				listRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(nil, domain.ErrInernalServerError)
			},
			res: domain.List{},
			err: domain.ErrInernalServerError,
		},
		{
			name: "success result",
			mock: func() {
				listRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(domain.List{}, nil)
			},
			res: domain.List{},
			err: nil,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.mock()

			res, err := listUseCase.GetHeadByID(context.Background(), 1)
			assert.Equal(t, tc.res, res)
			assert.ErrorIs(t, tc.err, err)
		})
	}
}

func TestGetPages(t *testing.T) {
	t.Parallel()

	listUseCase, _, pageRepo := list(t)
	tests := []test{
		{
			name: "error result",
			mock: func() {
				pageRepo.EXPECT().GetAll(gomock.Any()).Return(nil, domain.ErrInernalServerError)
			},
			res: []domain.Page(nil),
			err: domain.ErrInernalServerError,
		},
		{
			name: "success result",
			mock: func() {
				pageRepo.EXPECT().GetAll(gomock.Any()).Return([]domain.Page{}, nil)
			},
			res: []domain.Page{},
			err: nil,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.mock()
			res, err := listUseCase.GetPages(context.Background())
			assert.Equal(t, tc.res, res)
			assert.ErrorIs(t, tc.err, err)
		})
	}
}

func TestGetPageByID(t *testing.T) {
	t.Parallel()

	listUseCase, _, pageRepo := list(t)
	tests := []test{
		{
			name: "error result",
			mock: func() {
				pageRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(nil, domain.ErrInernalServerError)
			},
			res: []domain.Page(nil),
			err: domain.ErrInernalServerError,
		},
		{
			name: "success result",
			mock: func() {
				pageRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return([]domain.Page{}, nil)
			},
			res: []domain.Page{},
			err: nil,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.mock()
			res, err := listUseCase.GetPageByID(context.Background(), 1)
			assert.Equal(t, tc.res, res)
			assert.ErrorIs(t, tc.err, err)
		})
	}
}
