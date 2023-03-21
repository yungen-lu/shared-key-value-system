package integrationtest_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/charmbracelet/log"
	"github.com/docker/go-connections/nat"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	// _ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	v1 "github.com/yungen-lu/shared-key-value-list-system/internal/controller/http/v1"
	"github.com/yungen-lu/shared-key-value-list-system/internal/domain"
	"github.com/yungen-lu/shared-key-value-list-system/internal/usecase"
	"github.com/yungen-lu/shared-key-value-list-system/internal/usecase/repo"
	"github.com/yungen-lu/shared-key-value-list-system/pkg/postgres"
)

func startContainer(ctx context.Context) (testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
		Image:        "postgres:15-alpine",
		ExposedPorts: []string{"5432/tcp"},
		// WaitingFor:   wait.ForLog("database system is ready to accept connections"),
		WaitingFor: wait.ForAll(wait.ForLog("database system is ready to accept connections"), wait.ForListeningPort("5432/tcp")),
		Env: map[string]string{
			"POSTGRES_USER":     "postgres",
			"POSTGRES_PASSWORD": "postgres",
		},
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}
	return container, err
}
func startPostgres(host string, port string) (*postgres.Postgres, error) {
	pg, err := postgres.New(fmt.Sprintf("postgres://postgres:postgres@%s:%s?sslmode=disable", host, port))
	return pg, err
}

func migrateUp(host string, port string) error {
	m, err := migrate.New("file://../db/migrations", fmt.Sprintf("postgres://postgres:postgres@%s:%s?sslmode=disable", host, port))
	if err != nil {
		return err
	}
	defer m.Close()
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
func migrateDown(host string, port string) error {
	m, err := migrate.New("file://../db/migrations", fmt.Sprintf("postgres://postgres:postgres@%s:%s?sslmode=disable", host, port))
	if err != nil {
		return err
	}
	defer m.Close()
	if err = m.Down(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

var (
	containerHost string
	containerPort string
)

func TestMain(m *testing.M) {
	ctx := context.Background()

	port, err := nat.NewPort("tcp", "5432")
	if err != nil {
		log.Fatal("can't create port", "err", err)
	}
	container, err := startContainer(ctx)
	if err != nil {
		log.Fatal("can't start container", "err", err)
	}
	p, err := container.MappedPort(ctx, port)
	if err != nil {
		log.Fatal("can't get mapped port", "err", err)
	}
	containerPort = p.Port()
	containerHost, err = container.Host(ctx)
	if err != nil {
		log.Fatal("can't get container host", "err", err)
	}
	// err = migrateUp(containerHost, containerPort)
	// if err != nil {
	// 	log.Fatal("can't migrate up", "err", err)
	// }

	code := m.Run()

	if err := container.Terminate(ctx); err != nil {
		log.Error("err when terminating container", "err", err)
	}
	os.Exit(code)
}

type test struct {
	name    string
	url     string
	method  string
	payload interface{}
	code    int
	res     interface{}
}

func TestHead(t *testing.T) {
	pg, err := startPostgres(containerHost, containerPort)
	assert.NoError(t, err)
	defer pg.Close()

	err = migrateUp(containerHost, containerPort)
	assert.NoError(t, err)
	defer migrateDown(containerHost, containerPort)

	router := gin.Default()

	list := usecase.NewListUseCase(repo.NewListRepo(pg.Pool), repo.NewPageRepo(pg.Pool), 5*time.Second)
	v1.NewRouter(router, list)
	tests := []test{
		{
			name:   "add head",
			url:    "/v1/head",
			method: "POST",
			payload: v1.CreateHeadRequest{
				Key: "test-head",
			},
			code: http.StatusOK,
			res:  nil,
		},
		{
			name:    "get head",
			url:     "/v1/head/test-head",
			method:  "GET",
			payload: nil,
			code:    http.StatusOK,
			res: domain.List{
				ID:          1,
				Key:         "test-head",
				NextPageKey: nil,
			},
		},
		{
			name:   "add head 2",
			url:    "/v1/head",
			method: "POST",
			payload: v1.CreateHeadRequest{
				Key:         "test-head2",
				NextPageKey: nil,
			},
			code: http.StatusOK,
			res:  nil,
		},
		{
			name:    "get heads",
			url:     "/v1/head",
			method:  "GET",
			payload: nil,
			code:    http.StatusOK,
			res: []domain.List{
				{
					ID:          1,
					Key:         "test-head",
					NextPageKey: nil,
				},
				{
					ID:          2,
					Key:         "test-head2",
					NextPageKey: nil,
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			p, _ := json.Marshal(tc.payload)
			router.ServeHTTP(recorder, httptest.NewRequest(tc.method, tc.url, bytes.NewBuffer(p)))
			assert.Equal(t, tc.code, recorder.Code)
			if tc.res != nil {
				res, _ := json.Marshal(tc.res)
				assert.Equal(t, string(res), recorder.Body.String())
			}
		})
	}
}

func TestInsertPageAtTail(t *testing.T) {
	pg, err := startPostgres(containerHost, containerPort)
	assert.NoError(t, err)
	defer pg.Close()

	err = migrateUp(containerHost, containerPort)
	assert.NoError(t, err)
	defer migrateDown(containerHost, containerPort)

	router := gin.Default()

	list := usecase.NewListUseCase(repo.NewListRepo(pg.Pool), repo.NewPageRepo(pg.Pool), 5*time.Second)
	v1.NewRouter(router, list)
	tests := []test{
		{
			name:   "add head",
			url:    "/v1/head",
			method: "POST",
			payload: v1.CreateHeadRequest{
				Key: "test-head",
			},
			code: http.StatusOK,
			res:  nil,
		},
		{
			name:   "add page",
			url:    "/v1/page",
			method: "POST",
			payload: v1.CreatePageRequest{
				Key:     "test-page",
				ListKey: "test-head",
			},
			code: http.StatusOK,
			res:  nil,
		},
		{
			name:    "get page",
			url:     "/v1/page/test-page",
			method:  "GET",
			payload: nil,
			code:    http.StatusOK,
			res: domain.Page{
				ID:          1,
				Key:         "test-page",
				Articles:    nil,
				NextPageKey: nil,
				ListKey:     "test-head",
			},
		},
		{
			name:   "add page 2",
			url:    "/v1/page",
			method: "POST",
			payload: v1.CreatePageRequest{
				Key:     "test-page2",
				ListKey: "test-head",
			},
			code: http.StatusOK,
			res:  nil,
		},
		{
			name:    "get pages",
			url:     "/v1/page",
			method:  "GET",
			payload: nil,
			code:    http.StatusOK,
			res: []domain.Page{
				{
					ID:          1,
					Key:         "test-page",
					Articles:    nil,
					NextPageKey: ptr("test-page2"),
					ListKey:     "test-head",
				},
				{
					ID:          2,
					Key:         "test-page2",
					Articles:    nil,
					NextPageKey: nil,
					ListKey:     "test-head",
				},
			},
		},
		{
			name:    "get head",
			url:     "/v1/head/test-head",
			method:  "GET",
			payload: nil,
			code:    http.StatusOK,
			res: domain.List{
				ID:            1,
				Key:           "test-head",
				NextPageKey:   ptr("test-page"),
				LatestPageKey: ptr("test-page2"),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			p, _ := json.Marshal(tc.payload)
			router.ServeHTTP(recorder, httptest.NewRequest(tc.method, tc.url, bytes.NewBuffer(p)))
			assert.Equal(t, tc.code, recorder.Code)
			if tc.res != nil {
				res, _ := json.Marshal(tc.res)
				assert.Equal(t, string(res), recorder.Body.String())
			}
		})
	}
}

func TestInsertPageAtHead(t *testing.T) {
	pg, err := startPostgres(containerHost, containerPort)
	assert.NoError(t, err)
	defer pg.Close()

	err = migrateUp(containerHost, containerPort)
	assert.NoError(t, err)
	defer migrateDown(containerHost, containerPort)

	router := gin.Default()

	list := usecase.NewListUseCase(repo.NewListRepo(pg.Pool), repo.NewPageRepo(pg.Pool), 5*time.Second)
	v1.NewRouter(router, list)
	tests := []test{
		{
			name:   "add head",
			url:    "/v1/head",
			method: "POST",
			payload: v1.CreateHeadRequest{
				Key: "test-head",
			},
			code: http.StatusOK,
			res:  nil,
		},
		{
			name:   "add page",
			url:    "/v1/page",
			method: "POST",
			payload: v1.CreatePageRequest{
				Key:     "test-page",
				ListKey: "test-head",
			},
			code: http.StatusOK,
			res:  nil,
		},
		{
			name:    "get page",
			url:     "/v1/page/test-page",
			method:  "GET",
			payload: nil,
			code:    http.StatusOK,
			res: domain.Page{
				ID:          1,
				Key:         "test-page",
				Articles:    nil,
				NextPageKey: nil,
				ListKey:     "test-head",
			},
		},
		{
			name:   "add page 2",
			url:    "/v1/page",
			method: "POST",
			payload: v1.CreatePageRequest{
				Key:         "test-page2",
				NextPageKey: ptr("test-page"),
				ListKey:     "test-head",
			},
			code: http.StatusOK,
			res:  nil,
		},
		{
			name:    "get pages",
			url:     "/v1/page",
			method:  "GET",
			payload: nil,
			code:    http.StatusOK,
			res: []domain.Page{
				{
					ID:          1,
					Key:         "test-page",
					Articles:    nil,
					NextPageKey: nil,
					ListKey:     "test-head",
				},
				{
					ID:          2,
					Key:         "test-page2",
					Articles:    nil,
					NextPageKey: ptr("test-page"),
					ListKey:     "test-head",
				},
			},
		},
		{
			name:    "get head",
			url:     "/v1/head/test-head",
			method:  "GET",
			payload: nil,
			code:    http.StatusOK,
			res: domain.List{
				ID:            1,
				Key:           "test-head",
				NextPageKey:   ptr("test-page2"),
				LatestPageKey: ptr("test-page"),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			p, _ := json.Marshal(tc.payload)
			router.ServeHTTP(recorder, httptest.NewRequest(tc.method, tc.url, bytes.NewBuffer(p)))
			assert.Equal(t, tc.code, recorder.Code)
			if tc.res != nil {
				res, _ := json.Marshal(tc.res)
				assert.Equal(t, string(res), recorder.Body.String())
			}
		})
	}
}

func TestInsertPageAtMid(t *testing.T) {
	pg, err := startPostgres(containerHost, containerPort)
	assert.NoError(t, err)
	defer pg.Close()

	err = migrateUp(containerHost, containerPort)
	assert.NoError(t, err)
	defer migrateDown(containerHost, containerPort)

	router := gin.Default()

	list := usecase.NewListUseCase(repo.NewListRepo(pg.Pool), repo.NewPageRepo(pg.Pool), 5*time.Second)
	v1.NewRouter(router, list)
	tests := []test{
		{
			name:   "add head",
			url:    "/v1/head",
			method: "POST",
			payload: v1.CreateHeadRequest{
				Key: "test-head",
			},
			code: http.StatusOK,
			res:  nil,
		},
		{
			name:   "add page",
			url:    "/v1/page",
			method: "POST",
			payload: v1.CreatePageRequest{
				Key:     "test-page",
				ListKey: "test-head",
			},
			code: http.StatusOK,
			res:  nil,
		},
		{
			name:   "add page 2",
			url:    "/v1/page",
			method: "POST",
			payload: v1.CreatePageRequest{
				Key:     "test-page2",
				ListKey: "test-head",
			},
			code: http.StatusOK,
			res:  nil,
		},
		{
			name:   "add page 3",
			url:    "/v1/page",
			method: "POST",
			payload: v1.CreatePageRequest{
				Key:         "test-page3",
				NextPageKey: ptr("test-page2"),
				ListKey:     "test-head",
			},
			code: http.StatusOK,
			res:  nil,
		},
		{
			name:    "get pages",
			url:     "/v1/page",
			method:  "GET",
			payload: nil,
			code:    http.StatusOK,
			res: []domain.Page{
				{
					ID:          1,
					Key:         "test-page",
					Articles:    nil,
					NextPageKey: ptr("test-page3"),
					ListKey:     "test-head",
				},
				{
					ID:          2,
					Key:         "test-page2",
					Articles:    nil,
					NextPageKey: nil,
					ListKey:     "test-head",
				},
				{
					ID:          3,
					Key:         "test-page3",
					Articles:    nil,
					NextPageKey: ptr("test-page2"),
					ListKey:     "test-head",
				},
			},
		},
		{
			name:    "get head",
			url:     "/v1/head/test-head",
			method:  "GET",
			payload: nil,
			code:    http.StatusOK,
			res: domain.List{
				ID:            1,
				Key:           "test-head",
				NextPageKey:   ptr("test-page"),
				LatestPageKey: ptr("test-page2"),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			p, _ := json.Marshal(tc.payload)
			router.ServeHTTP(recorder, httptest.NewRequest(tc.method, tc.url, bytes.NewBuffer(p)))
			assert.Equal(t, tc.code, recorder.Code)
			if tc.res != nil {
				res, _ := json.Marshal(tc.res)
				assert.Equal(t, string(res), recorder.Body.String())
			}
		})
	}
}
func TestDeletePageAtTail(t *testing.T) {
	pg, err := startPostgres(containerHost, containerPort)
	assert.NoError(t, err)
	defer pg.Close()

	err = migrateUp(containerHost, containerPort)
	assert.NoError(t, err)
	defer migrateDown(containerHost, containerPort)

	router := gin.Default()

	list := usecase.NewListUseCase(repo.NewListRepo(pg.Pool), repo.NewPageRepo(pg.Pool), 5*time.Second)
	v1.NewRouter(router, list)
	tests := []test{
		{
			name:   "add head",
			url:    "/v1/head",
			method: "POST",
			payload: v1.CreateHeadRequest{
				Key: "test-head",
			},
			code: http.StatusOK,
			res:  nil,
		},
		{
			name:   "add page",
			url:    "/v1/page",
			method: "POST",
			payload: v1.CreatePageRequest{
				Key:     "test-page",
				ListKey: "test-head",
			},
			code: http.StatusOK,
			res:  nil,
		},
		{
			name:   "add page 2",
			url:    "/v1/page",
			method: "POST",
			payload: v1.CreatePageRequest{
				Key:     "test-page2",
				ListKey: "test-head",
			},
			code: http.StatusOK,
			res:  nil,
		},
		{
			name:    "delete page at tail",
			url:     "/v1/page/test-page2",
			method:  "DELETE",
			payload: nil,
			code:    http.StatusOK,
			res:     nil,
		},
		{
			name:    "get pages",
			url:     "/v1/page",
			method:  "GET",
			payload: nil,
			code:    http.StatusOK,
			res: []domain.Page{
				{
					ID:          1,
					Key:         "test-page",
					Articles:    nil,
					NextPageKey: nil,
					ListKey:     "test-head",
				},
			},
		},
		{
			name:    "get head",
			url:     "/v1/head/test-head",
			method:  "GET",
			payload: nil,
			code:    http.StatusOK,
			res: domain.List{
				ID:            1,
				Key:           "test-head",
				NextPageKey:   ptr("test-page"),
				LatestPageKey: ptr("test-page"),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			p, _ := json.Marshal(tc.payload)
			router.ServeHTTP(recorder, httptest.NewRequest(tc.method, tc.url, bytes.NewBuffer(p)))
			assert.Equal(t, tc.code, recorder.Code)
			if tc.res != nil {
				res, _ := json.Marshal(tc.res)
				assert.Equal(t, string(res), recorder.Body.String())
			}
		})
	}
}

func TestDeletePageAtHead(t *testing.T) {
	pg, err := startPostgres(containerHost, containerPort)
	assert.NoError(t, err)
	defer pg.Close()

	err = migrateUp(containerHost, containerPort)
	assert.NoError(t, err)
	defer migrateDown(containerHost, containerPort)

	router := gin.Default()

	list := usecase.NewListUseCase(repo.NewListRepo(pg.Pool), repo.NewPageRepo(pg.Pool), 5*time.Second)
	v1.NewRouter(router, list)
	tests := []test{
		{
			name:   "add head",
			url:    "/v1/head",
			method: "POST",
			payload: v1.CreateHeadRequest{
				Key: "test-head",
			},
			code: http.StatusOK,
			res:  nil,
		},
		{
			name:   "add page",
			url:    "/v1/page",
			method: "POST",
			payload: v1.CreatePageRequest{
				Key:     "test-page",
				ListKey: "test-head",
			},
			code: http.StatusOK,
			res:  nil,
		},
		{
			name:   "add page 2",
			url:    "/v1/page",
			method: "POST",
			payload: v1.CreatePageRequest{
				Key:     "test-page2",
				ListKey: "test-head",
			},
			code: http.StatusOK,
			res:  nil,
		},
		{
			name:    "delete page at head",
			url:     "/v1/page/test-page",
			method:  "DELETE",
			payload: nil,
			code:    http.StatusOK,
			res:     nil,
		},
		{
			name:    "get pages",
			url:     "/v1/page",
			method:  "GET",
			payload: nil,
			code:    http.StatusOK,
			res: []domain.Page{
				{
					ID:          2,
					Key:         "test-page2",
					Articles:    nil,
					NextPageKey: nil,
					ListKey:     "test-head",
				},
			},
		},
		{
			name:    "get head",
			url:     "/v1/head/test-head",
			method:  "GET",
			payload: nil,
			code:    http.StatusOK,
			res: domain.List{
				ID:            1,
				Key:           "test-head",
				NextPageKey:   ptr("test-page2"),
				LatestPageKey: ptr("test-page2"),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			p, _ := json.Marshal(tc.payload)
			router.ServeHTTP(recorder, httptest.NewRequest(tc.method, tc.url, bytes.NewBuffer(p)))
			assert.Equal(t, tc.code, recorder.Code)
			if tc.res != nil {
				res, _ := json.Marshal(tc.res)
				assert.Equal(t, string(res), recorder.Body.String())
			}
		})
	}
}

func TestDeletePageAtMid(t *testing.T) {
	pg, err := startPostgres(containerHost, containerPort)
	assert.NoError(t, err)
	defer pg.Close()

	err = migrateUp(containerHost, containerPort)
	assert.NoError(t, err)
	defer migrateDown(containerHost, containerPort)

	router := gin.Default()

	list := usecase.NewListUseCase(repo.NewListRepo(pg.Pool), repo.NewPageRepo(pg.Pool), 5*time.Second)
	v1.NewRouter(router, list)
	tests := []test{
		{
			name:   "add head",
			url:    "/v1/head",
			method: "POST",
			payload: v1.CreateHeadRequest{
				Key: "test-head",
			},
			code: http.StatusOK,
			res:  nil,
		},
		{
			name:   "add page",
			url:    "/v1/page",
			method: "POST",
			payload: v1.CreatePageRequest{
				Key:     "test-page",
				ListKey: "test-head",
			},
			code: http.StatusOK,
			res:  nil,
		},
		{
			name:   "add page 2",
			url:    "/v1/page",
			method: "POST",
			payload: v1.CreatePageRequest{
				Key:     "test-page2",
				ListKey: "test-head",
			},
			code: http.StatusOK,
			res:  nil,
		},
		{
			name:   "add page 3",
			url:    "/v1/page",
			method: "POST",
			payload: v1.CreatePageRequest{
				Key:     "test-page3",
				ListKey: "test-head",
			},
			code: http.StatusOK,
			res:  nil,
		},
		{
			name:    "delete page at mid",
			url:     "/v1/page/test-page2",
			method:  "DELETE",
			payload: nil,
			code:    http.StatusOK,
			res:     nil,
		},
		{
			name:    "get pages",
			url:     "/v1/page",
			method:  "GET",
			payload: nil,
			code:    http.StatusOK,
			res: []domain.Page{
				{
					ID:          1,
					Key:         "test-page",
					Articles:    nil,
					NextPageKey: ptr("test-page3"),
					ListKey:     "test-head",
				},
				{
					ID:          3,
					Key:         "test-page3",
					Articles:    nil,
					NextPageKey: nil,
					ListKey:     "test-head",
				},
			},
		},
		{
			name:    "get head",
			url:     "/v1/head/test-head",
			method:  "GET",
			payload: nil,
			code:    http.StatusOK,
			res: domain.List{
				ID:            1,
				Key:           "test-head",
				NextPageKey:   ptr("test-page"),
				LatestPageKey: ptr("test-page3"),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			p, _ := json.Marshal(tc.payload)
			router.ServeHTTP(recorder, httptest.NewRequest(tc.method, tc.url, bytes.NewBuffer(p)))
			assert.Equal(t, tc.code, recorder.Code)
			if tc.res != nil {
				res, _ := json.Marshal(tc.res)
				assert.Equal(t, string(res), recorder.Body.String())
			}
		})
	}
}
func ptr[T any](v T) *T {
	return &v
}
