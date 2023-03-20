package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/yungen-lu/shared-key-value-list-system/docs"
	"github.com/yungen-lu/shared-key-value-list-system/internal/controller/http/middleware"
	"github.com/yungen-lu/shared-key-value-list-system/internal/usecase"
)

// @title		Shared Key Value List System API
// @version	1.0
// @basePath	/v1
func NewRouter(handler *gin.Engine, listUseCase usecase.List) {
	handler.Use(gin.Recovery())
	handler.Use(middleware.CustomLogger())
	handler.GET("/healthz", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})
	v1 := handler.Group("/v1")
	{
		NewHeadRoutes(v1, listUseCase)
		NewPageRoutes(v1, listUseCase)
	}
	handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
