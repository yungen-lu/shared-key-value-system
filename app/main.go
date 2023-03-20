package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/yungen-lu/shared-key-value-list-system/config"
	v1 "github.com/yungen-lu/shared-key-value-list-system/internal/controller/http/v1"
	"github.com/yungen-lu/shared-key-value-list-system/internal/usecase"
	"github.com/yungen-lu/shared-key-value-list-system/internal/usecase/repo"
	"github.com/yungen-lu/shared-key-value-list-system/pkg/httpserver"
	"github.com/yungen-lu/shared-key-value-list-system/pkg/postgres"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("can't load config", "err", err)
	}
	pg, err := postgres.New(cfg.URL)
	if err != nil {
		log.Fatal("can't connect to postgres", "err", err)
	}
	defer pg.Close()

	handler := gin.New()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Info("[GIN-debug]", "method", httpMethod, "path", absolutePath, "handler", handlerName, "number of handlers", nuHandlers)
	}
	list := usecase.NewListUseCase(repo.NewListRepo(pg.Pool), repo.NewPageRepo(pg.Pool), 5*time.Second)
	v1.NewRouter(handler, list)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case sig := <-interrupt:
		log.Info("received signal", "signal", sig.String())
	case err := <-httpServer.Notify():
		log.Error("http server error", "err", err)
	}
	err = httpServer.Shutdown()
	if err != nil {
		log.Error("err when shutting down http server", "err", err)
	}
}
