package v1_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	v1 "github.com/yungen-lu/shared-key-value-list-system/internal/controller/http/v1"
	"github.com/yungen-lu/shared-key-value-list-system/internal/domain"
)

func TestGetPages(t *testing.T) {
	listUseCase := list(t)
	router := gin.Default()
	v1.NewPageRoutes(router.Group("/v1"), listUseCase)

	tests := []test{
		{
			name: "success result",
			mock: func() {
				listUseCase.EXPECT().GetPages(gomock.Any()).Return([]domain.Page{}, nil)
			},
			code: http.StatusOK,
			res:  []domain.Page{},
		},
		{
			name: "internal error result",
			mock: func() {
				listUseCase.EXPECT().GetPages(gomock.Any()).Return(nil, domain.ErrInernalServerError)
			},
			code: http.StatusInternalServerError,
			res:  nil,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, httptest.NewRequest("GET", "/v1/page", nil))
			assert.Equal(t, tc.code, recorder.Code)
			if tc.res != nil {
				res, _ := json.Marshal(tc.res)
				assert.Equal(t, string(res), recorder.Body.String())
			}
		})
	}
}

func TestGetPageByKey(t *testing.T) {
	listUseCase := list(t)
	router := gin.Default()
	v1.NewPageRoutes(router.Group("/v1"), listUseCase)

	tests := []test{
		{
			name:  "success result",
			param: "test-page",
			mock: func() {
				listUseCase.EXPECT().GetPageByKey(gomock.Any(), gomock.Eq("test-page")).Return(domain.Page{}, nil)
			},
			code: http.StatusOK,
			res:  domain.Page{},
		},
		{
			name:  "bad param",
			param: "not-exists",
			mock: func() {
				listUseCase.EXPECT().GetPageByKey(gomock.Any(), gomock.Eq("not-exists")).Return(domain.Page{}, domain.ErrInernalServerError)
			},
			code: http.StatusInternalServerError,
			res:  nil,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, httptest.NewRequest("GET", "/v1/page/"+tc.param, nil))
			assert.Equal(t, tc.code, recorder.Code)
			if tc.res != nil {
				res, _ := json.Marshal(tc.res)
				assert.Equal(t, string(res), recorder.Body.String())
			}
		})
	}
}

func TestCreatePage(t *testing.T) {
	listUseCase := list(t)
	router := gin.Default()
	v1.NewPageRoutes(router.Group("/v1"), listUseCase)

	tests := []test{
		{
			name: "success result",
			mock: func() {
				listUseCase.EXPECT().CreatePage(gomock.Any(), domain.Page{Key: "test-page", ListKey: "test-list"}).Return(nil)
			},
			payload: v1.CreatePageRequest{Key: "test-page", ListKey: "test-list"},
			code:    http.StatusOK,
			res:     nil,
		},
		{
			name: "null body",
			mock: func() {
			},
			payload: nil,
			code:    http.StatusBadRequest,
			res:     nil,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			recorder := httptest.NewRecorder()
			p, _ := json.Marshal(tc.payload)
			router.ServeHTTP(recorder, httptest.NewRequest("POST", "/v1/page", bytes.NewBuffer(p)))
			assert.Equal(t, tc.code, recorder.Code)
			if tc.res != nil {
				res, _ := json.Marshal(tc.res)
				assert.Equal(t, string(res), recorder.Body.String())
			}
		})
	}
}

func TestUpdatePageByKey(t *testing.T) {
	listUseCase := list(t)
	router := gin.Default()
	v1.NewPageRoutes(router.Group("/v1"), listUseCase)
	next_page := "next-page"
	tests := []test{
		{
			name:  "success result",
			param: "test-page",
			mock: func() {
				listUseCase.EXPECT().UpdatePageByKey(gomock.Any(), gomock.Eq("test-page"), domain.Page{NextPageKey: &next_page}).Return(nil)
			},
			payload: v1.UpdatePageRequest{NextPageKey: &next_page},
			code:    http.StatusOK,
			res:     nil,
		},
		{
			name:  "failed result",
			param: "test-page",
			mock: func() {
				listUseCase.EXPECT().UpdatePageByKey(gomock.Any(), gomock.Eq("test-page"), domain.Page{NextPageKey: nil}).Return(domain.ErrInernalServerError)
			},
			payload: nil,
			code:    http.StatusInternalServerError,
			res:     nil,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			recorder := httptest.NewRecorder()
			p, _ := json.Marshal(tc.payload)
			router.ServeHTTP(recorder, httptest.NewRequest("PUT", "/v1/page/"+tc.param, bytes.NewBuffer(p)))
			assert.Equal(t, tc.code, recorder.Code)
			if tc.res != nil {
				res, _ := json.Marshal(tc.res)
				assert.Equal(t, string(res), recorder.Body.String())
			}
		})
	}
}
