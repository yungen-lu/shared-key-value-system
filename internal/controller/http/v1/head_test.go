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
	"github.com/yungen-lu/shared-key-value-list-system/internal/usecase/mocks"
)

type test struct {
	name    string
	mock    func()
	payload interface{}
	param   string
	code    int
	res     interface{}
}

func list(t *testing.T) *mocks.MockList {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	return mocks.NewMockList(ctrl)
}
func TestGetHeads(t *testing.T) {
	listUseCase := list(t)
	router := gin.Default()
	v1.NewHeadRoutes(router.Group("/v1"), listUseCase)

	tests := []test{
		{
			name: "success result",
			mock: func() {
				listUseCase.EXPECT().GetHeads(gomock.Any()).Return([]domain.List{}, nil)
			},
			code: http.StatusOK,
			res:  []domain.List{},
		},
		{
			name: "internal error result",
			mock: func() {
				listUseCase.EXPECT().GetHeads(gomock.Any()).Return(nil, domain.ErrInernalServerError)
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
			router.ServeHTTP(recorder, httptest.NewRequest("GET", "/v1/head", nil))
			assert.Equal(t, tc.code, recorder.Code)
			if tc.res != nil {
				res, _ := json.Marshal(tc.res)
				assert.Equal(t, string(res), recorder.Body.String())
			}
		})
	}
}

func TestGetHeadByKey(t *testing.T) {
	listUseCase := list(t)
	router := gin.Default()
	v1.NewHeadRoutes(router.Group("/v1"), listUseCase)

	tests := []test{
		{
			name:  "success result",
			param: "test-head",
			mock: func() {
				listUseCase.EXPECT().GetHeadByKey(gomock.Any(), gomock.Eq("test-head")).Return(domain.List{}, nil)
			},
			code: http.StatusOK,
			res:  domain.List{},
		},
		{
			name:  "bad param",
			param: "not-exists",
			mock: func() {
				listUseCase.EXPECT().GetHeadByKey(gomock.Any(), gomock.Eq("not-exists")).Return(domain.List{}, domain.ErrInernalServerError)
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
			router.ServeHTTP(recorder, httptest.NewRequest("GET", "/v1/head/"+tc.param, nil))
			assert.Equal(t, tc.code, recorder.Code)
			if tc.res != nil {
				res, _ := json.Marshal(tc.res)
				assert.Equal(t, string(res), recorder.Body.String())
			}
		})
	}
}

func TestCreateHead(t *testing.T) {
	listUseCase := list(t)
	router := gin.Default()
	v1.NewHeadRoutes(router.Group("/v1"), listUseCase)

	tests := []test{
		{
			name: "success result",
			mock: func() {
				listUseCase.EXPECT().CreateHead(gomock.Any(), domain.List{Key: "test-head"}).Return(nil)
			},
			payload: v1.CreateHeadRequest{Key: "test-head"},
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
			router.ServeHTTP(recorder, httptest.NewRequest("POST", "/v1/head", bytes.NewBuffer(p)))
			assert.Equal(t, tc.code, recorder.Code)
			if tc.res != nil {
				res, _ := json.Marshal(tc.res)
				assert.Equal(t, string(res), recorder.Body.String())
			}
		})
	}
}

func TestUpdateHeadByKey(t *testing.T) {
	listUseCase := list(t)
	router := gin.Default()
	v1.NewHeadRoutes(router.Group("/v1"), listUseCase)
	next_head := "next-head"
	tests := []test{
		{
			name:  "success result",
			param: "test-head",
			mock: func() {
				listUseCase.EXPECT().UpdateHeadByKey(gomock.Any(), gomock.Eq("test-head"), domain.List{NextPageKey: &next_head}).Return(nil)
			},
			payload: v1.UpdateHeadRequest{NextPageKey: &next_head},
			code:    http.StatusOK,
			res:     nil,
		},
		{
			name:  "failed result",
			param: "test-head",
			mock: func() {
				listUseCase.EXPECT().UpdateHeadByKey(gomock.Any(), gomock.Eq("test-head"), domain.List{NextPageKey: nil}).Return(domain.ErrInernalServerError)
			},
			payload: nil,
			code:    http.StatusInternalServerError,
			// code:    http.StatusBadRequest,
			res: nil,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			recorder := httptest.NewRecorder()
			p, _ := json.Marshal(tc.payload)
			router.ServeHTTP(recorder, httptest.NewRequest("PUT", "/v1/head/"+tc.param, bytes.NewBuffer(p)))
			assert.Equal(t, tc.code, recorder.Code)
			if tc.res != nil {
				res, _ := json.Marshal(tc.res)
				assert.Equal(t, string(res), recorder.Body.String())
			}
		})
	}
}

func TestDeleteHead(t *testing.T) {
	listUseCase := list(t)
	router := gin.Default()
	v1.NewHeadRoutes(router.Group("/v1"), listUseCase)

	tests := []test{
		{
			name:  "success result",
			param: "/test-head",
			mock: func() {
				listUseCase.EXPECT().DeleteHeadByKey(gomock.Any(), gomock.Eq("test-head")).Return(nil)
			},
			payload: nil,
			code:    http.StatusOK,
			res:     nil,
		},
		{
			name:  "failed result",
			param: "/test-head",
			mock: func() {
				listUseCase.EXPECT().DeleteHeadByKey(gomock.Any(), gomock.Eq("test-head")).Return(domain.ErrInernalServerError)
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
			router.ServeHTTP(recorder, httptest.NewRequest("DELETE", "/v1/head"+tc.param, bytes.NewBuffer(p)))
			assert.Equal(t, tc.code, recorder.Code)
			if tc.res != nil {
				res, _ := json.Marshal(tc.res)
				assert.Equal(t, string(res), recorder.Body.String())
			}
		})
	}
}
