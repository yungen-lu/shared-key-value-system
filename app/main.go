package main

import (
	"os"
	"os/signal"
	"syscall"

	v1 "github.com/yungen-lu/shared-key-value-list-system/internal/controller/http/v1"
	"github.com/yungen-lu/shared-key-value-list-system/pkg/httpserver"
)

func main() {

	v1.NewRouter()
	httpServer := httpserver.New(v1.NewRouter(), httpserver.Port("80"))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case sig := <-interrupt:
		_ = sig
	case err := <-httpServer.Notify():
		_ = err
	}
}
