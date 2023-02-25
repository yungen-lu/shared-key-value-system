package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine) {
	handler.Use(gin.Recovery())
	handler.GET("/healthz", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})
	v1 := handler.Group("/v1")
	{
		v1.GET("/page", PageHandler)
	}
}
