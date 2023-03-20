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
	// p, err := filepath.Abs("./db/migrations")
	// if err != nil {
	// 	return err
	// }
	// log.Info("migrate path", "path", p)
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
	err = migrateUp(containerHost, containerPort)
	if err != nil {
		log.Fatal("can't migrate up", "err", err)
	}

	code := m.Run()

	if err := container.Terminate(ctx); err != nil {
		log.Error("err when terminating container", "err", err)
	}
	os.Exit(code)
}

type test struct {
	name    string
	param   string
	method  string
	payload interface{}
	code    int
	res     interface{}
}

func TestHead(t *testing.T) {
	pg, err := startPostgres(containerHost, containerPort)
	log.Info("postgres", "host", containerHost, "port", containerPort)
	assert.NoError(t, err)
	defer pg.Close()
	router := gin.Default()

	list := usecase.NewListUseCase(repo.NewListRepo(pg.Pool), repo.NewPageRepo(pg.Pool), 5*time.Second)
	v1.NewRouter(router, list)
	tests := []test{
		{
			name:   "add head",
			param:  "",
			method: "POST",
			payload: v1.CreateHeadRequest{
				Key: "test-head",
			},
			code: http.StatusOK,
			res:  nil,
		},
		{
			name:    "get head",
			param:   "/test-head",
			method:  "GET",
			payload: nil,
			code:    http.StatusOK,
			res: domain.List{
				ID:          1,
				Key:         "test-head",
				PageCount:   0,
				NextPageKey: nil,
			},
		},
		{
			name:   "add head 2",
			param:  "",
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
			param:   "",
			method:  "GET",
			payload: nil,
			code:    http.StatusOK,
			res: []domain.List{
				{
					ID:          1,
					Key:         "test-head",
					PageCount:   0,
					NextPageKey: nil,
				},
				{
					ID:          2,
					Key:         "test-head2",
					PageCount:   0,
					NextPageKey: nil,
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			p, _ := json.Marshal(tc.payload)
			router.ServeHTTP(recorder, httptest.NewRequest(tc.method, "/v1/head"+tc.param, bytes.NewBuffer(p)))
			assert.Equal(t, tc.code, recorder.Code)
			if tc.res != nil {
				res, _ := json.Marshal(tc.res)
				assert.Equal(t, string(res), recorder.Body.String())
			}
		})
	}
}

func TestPage(t *testing.T) {
	pg, err := startPostgres(containerHost, containerPort)
	log.Info("postgres", "host", containerHost, "port", containerPort)
	assert.NoError(t, err)
	defer pg.Close()
	router := gin.Default()

	list := usecase.NewListUseCase(repo.NewListRepo(pg.Pool), repo.NewPageRepo(pg.Pool), 5*time.Second)
	v1.NewRouter(router, list)
	tmp := "test-page"
	tests := []test{
		{
			name:   "add page",
			param:  "",
			method: "POST",
			payload: v1.CreatePageRequest{
				Key: "test-page",
			},
			code: http.StatusOK,
			res:  nil,
		},
		{
			name:    "get page",
			param:   "/test-page",
			method:  "GET",
			payload: nil,
			code:    http.StatusOK,
			res: domain.Page{
				ID:          1,
				Key:         "test-page",
				Articles:    nil,
				NextPageKey: nil,
			},
		},
		{
			name:   "add page 2",
			param:  "",
			method: "POST",
			payload: v1.CreatePageRequest{
				Key:         "test-page2",
				NextPageKey: &tmp,
			},
			code: http.StatusOK,
			res:  nil,
		},
		{
			name:    "get pages",
			param:   "",
			method:  "GET",
			payload: nil,
			code:    http.StatusOK,
			res: []domain.Page{
				{
					ID:          1,
					Key:         "test-page",
					Articles:    nil,
					NextPageKey: nil,
				},
				{
					ID:          2,
					Key:         "test-page2",
					Articles:    nil,
					NextPageKey: &tmp,
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			p, _ := json.Marshal(tc.payload)
			router.ServeHTTP(recorder, httptest.NewRequest(tc.method, "/v1/page"+tc.param, bytes.NewBuffer(p)))
			assert.Equal(t, tc.code, recorder.Code)
			if tc.res != nil {
				res, _ := json.Marshal(tc.res)
				assert.Equal(t, string(res), recorder.Body.String())
			}
		})
	}
}