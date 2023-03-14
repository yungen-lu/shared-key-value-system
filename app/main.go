package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	v1 "github.com/yungen-lu/shared-key-value-list-system/internal/controller/http/v1"
	"github.com/yungen-lu/shared-key-value-list-system/internal/usecase"
	"github.com/yungen-lu/shared-key-value-list-system/internal/usecase/repo"
	"github.com/yungen-lu/shared-key-value-list-system/pkg/httpserver"
	"github.com/yungen-lu/shared-key-value-list-system/pkg/postgres"
)

func main() {
	pg, err := postgres.New("postgresql://postgres@localhost:5432")
	if err != nil {
		panic(err)
	}
	defer pg.Close()

	handler := gin.New()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Info("[GIN-debug]", "method", httpMethod, "path", absolutePath, "handler", handlerName, "number of handlers", nuHandlers)
	}
	list := usecase.NewListUseCase(repo.NewListRepo(pg.Pool), repo.NewPageRepo(pg.Pool), 5*time.Second)
	v1.NewRouter(handler, list)
	httpServer := httpserver.New(handler, httpserver.Port("80"))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case sig := <-interrupt:
		_ = sig
	case err := <-httpServer.Notify():
		_ = err
	}
	err = httpServer.Shutdown()
	if err != nil {

	}
}
