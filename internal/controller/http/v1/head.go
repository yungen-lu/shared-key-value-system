package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yungen-lu/shared-key-value-list-system/internal/usecase"
)

type headRoutes struct {
	list usecase.List
}

func newHeadRoutes(handler *gin.RouterGroup) {
	r := &headRoutes{}
	h := handler.Group("/head")
	{
		h.GET("/", r.getAll)
		h.GET("/:id", r.getByID)
	}
}

//	@Summary		Get all heads
//	@Description	test
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	domain.List
//	@Falure			500 {object}
//	@Router			/head [get]
func (r *headRoutes) getAll(c *gin.Context) {
	lists, err := r.list.GetHeads(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, lists)
}

//	@Summary		Get a head by id
//	@Description	test
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Head ID"
//	@Success		200	{object}	domain.List
//	@Falure			500 {object}
//	@Router			/head/{id} [get]
func (r *headRoutes) getByID(c *gin.Context) {
	idStr, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	list, err := r.list.GetHeadByID(c, int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, list)
}
