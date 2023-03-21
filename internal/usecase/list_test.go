package usecase_test

import (
	"context"
	"testing"
	"time"

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
	list := usecase.NewListUseCase(listRepo, pageRepo, 5*time.Second)
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

func TestGetPageByKey(t *testing.T) {
	t.Parallel()
	listUsecase, _, pageRepo := list(t)
	tests := []test{
		{
			name: "error result",
			mock: func() {
				pageRepo.EXPECT().GetByKey(gomock.Any(), gomock.Eq("key")).Return(domain.Page{}, domain.ErrInernalServerError)
			},
			res: domain.Page{},
			err: domain.ErrInernalServerError,
		},
		{
			name: "success result",
			mock: func() {
				pageRepo.EXPECT().GetByKey(gomock.Any(), gomock.Eq("key")).Return(domain.Page{}, nil)
			},
			res: domain.Page{},
			err: nil,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.mock()
			res, err := listUsecase.GetPageByKey(context.Background(), "key")
			assert.Equal(t, tc.res, res)
			assert.ErrorIs(t, tc.err, err)
		})
	}
}

func TestGetHeadByKey(t *testing.T) {
	t.Parallel()

	listUsecase, listRepo, _ := list(t)
	tests := []test{
		{
			name: "error result",
			mock: func() {
				listRepo.EXPECT().GetByKey(gomock.Any(), gomock.Eq("key")).Return(domain.List{}, domain.ErrInernalServerError)
			},
			res: domain.List{},
			err: domain.ErrInernalServerError,
		},
		{
			name: "success result",
			mock: func() {
				listRepo.EXPECT().GetByKey(gomock.Any(), gomock.Eq("key")).Return(domain.List{}, nil)
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
			res, err := listUsecase.GetHeadByKey(context.Background(), "key")
			assert.Equal(t, tc.res, res)
			assert.ErrorIs(t, tc.err, err)
		})
	}
}

func TestCreateHead(t *testing.T) {
	t.Parallel()

	listUsecase, listRepo, _ := list(t)
	tests := []test{
		{
			name: "error result",
			mock: func() {
				listRepo.EXPECT().Store(gomock.Any(), domain.List{}).Return(domain.ErrInernalServerError)
			},
			err: domain.ErrInernalServerError,
		},
		{
			name: "success result",
			mock: func() {
				listRepo.EXPECT().Store(gomock.Any(), domain.List{}).Return(nil)
			},
			err: nil,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.mock()
			err := listUsecase.CreateHead(context.Background(), domain.List{})
			assert.ErrorIs(t, tc.err, err)
		})
	}
}

func TestCreatePage(t *testing.T) {
	t.Parallel()

	listUsecase, _, pageRepo := list(t)
	tests := []test{
		{
			name: "error result",
			mock: func() {
				pageRepo.EXPECT().Store(gomock.Any(), domain.Page{}).Return(domain.ErrInernalServerError)
			},
			err: domain.ErrInernalServerError,
		},
		{
			name: "success result",
			mock: func() {
				pageRepo.EXPECT().Store(gomock.Any(), domain.Page{}).Return(nil)
			},
			err: nil,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.mock()
			err := listUsecase.CreatePage(context.Background(), domain.Page{})
			assert.ErrorIs(t, tc.err, err)
		})
	}
}

func TestUpdateHeadByKey(t *testing.T) {
	t.Parallel()

	listUsecase, listRepo, _ := list(t)
	tests := []test{
		{
			name: "error result",
			mock: func() {
				listRepo.EXPECT().UpdateByKey(gomock.Any(), gomock.Eq("key"), domain.List{}).Return(domain.ErrInernalServerError)
			},
			err: domain.ErrInernalServerError,
		},
		{
			name: "success result",
			mock: func() {
				listRepo.EXPECT().UpdateByKey(gomock.Any(), gomock.Eq("key"), domain.List{}).Return(nil)
			},
			err: nil,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.mock()
			err := listUsecase.UpdateHeadByKey(context.Background(), "key", domain.List{})
			assert.ErrorIs(t, tc.err, err)
		})
	}
}

func TestDeleteHeadByKey(t *testing.T) {
	t.Parallel()

	listUsecase, listRepo, _ := list(t)
	tests := []test{
		{
			name: "error result",
			mock: func() {
				listRepo.EXPECT().DeleteByKey(gomock.Any(), gomock.Eq("key")).Return(domain.ErrInernalServerError)
			},
			err: domain.ErrInernalServerError,
		},
		{
			name: "success result",
			mock: func() {
				listRepo.EXPECT().DeleteByKey(gomock.Any(), gomock.Eq("key")).Return(nil)
			},
			err: nil,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.mock()
			err := listUsecase.DeleteHeadByKey(context.Background(), "key")
			assert.ErrorIs(t, tc.err, err)
		})
	}
}

func TestDeletePageByKey(t *testing.T) {
	t.Parallel()

	listUsecase, _, pageRepo := list(t)
	tests := []test{
		{
			name: "error result",
			mock: func() {
				pageRepo.EXPECT().DeleteByKey(gomock.Any(), gomock.Eq("key")).Return(domain.ErrInernalServerError)
			},
			err: domain.ErrInernalServerError,
		},
		{
			name: "success result",
			mock: func() {
				pageRepo.EXPECT().DeleteByKey(gomock.Any(), gomock.Eq("key")).Return(nil)
			},
			err: nil,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.mock()
			err := listUsecase.DeletePageByKey(context.Background(), "key")
			assert.ErrorIs(t, tc.err, err)
		})
	}
}

func TestOutdatedLists(t *testing.T) {
	t.Parallel()

	listUsecase, listRepo, _ := list(t)
	tests := []test{
		{
			name: "error result",
			mock: func() {
				listRepo.EXPECT().DeleteOutdated(gomock.Any()).Return(int64(0), domain.ErrInernalServerError)
			},
			res: int64(0),
			err: domain.ErrInernalServerError,
		},
		{
			name: "success result",
			mock: func() {
				listRepo.EXPECT().DeleteOutdated(gomock.Any()).Return(int64(1), nil)
			},
			res: int64(1),
			err: nil,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.mock()
			count, err := listUsecase.DeleteOutdatedLists(context.Background())
			assert.Equal(t, tc.res, count)
			assert.ErrorIs(t, tc.err, err)
		})
	}
}
