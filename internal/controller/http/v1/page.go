package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PageHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}
