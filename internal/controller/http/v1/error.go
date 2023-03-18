package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yungen-lu/shared-key-value-list-system/internal/domain"
)

type response struct {
	Error string `json:"error"`
}

func responseError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, domain.ErrInernalServerError):
		c.JSON(http.StatusInternalServerError, response{Error: err.Error()})
	case errors.Is(err, domain.ErrBadParamInput):
		c.JSON(http.StatusBadRequest, response{Error: err.Error()})

	case errors.Is(err, domain.ErrNotFound):
		c.JSON(http.StatusNotFound, response{Error: err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, response{Error: err.Error()})
	}
}
