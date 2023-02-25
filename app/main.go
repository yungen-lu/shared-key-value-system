package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	v1 "github.com/yungen-lu/shared-key-value-list-system/internal/controller/http/v1"
	"github.com/yungen-lu/shared-key-value-list-system/pkg/httpserver"
)

func main() {

	handler := gin.New()
	v1.NewRouter(handler)
	httpServer := httpserver.New(handler, httpserver.Port("80"))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case sig := <-interrupt:
		_ = sig
	case err := <-httpServer.Notify():
		_ = err
	}
	err := httpServer.Shutdown()
	if err != nil {

	}
}
