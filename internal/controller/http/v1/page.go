package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yungen-lu/shared-key-value-list-system/internal/usecase"
)

type pageRoutes struct {
	list usecase.List
}

func newPageRoutes(handler *gin.RouterGroup) {
	r := &pageRoutes{}
	h := handler.Group("/page")
	{
		h.GET("/", r.getAll)
		h.GET("/:id", r.getByID)
	}
}

//	@Summary		Get all pages
//	@Description	test
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	domain.Page
//	@Falure			500 {object}
//	@Router			/page [get]
func (r *pageRoutes) getAll(c *gin.Context) {
	pages, err := r.list.GetPages(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, pages)
}

//	@Summary		Get a page by id
//	@Description	test
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"page ID"
//	@Success		200	{object}	domain.Page
//	@Falure			500 {object}
//	@Router			/page/{id} [get]
func (r *pageRoutes) getByID(c *gin.Context) {
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
	page, err := r.list.GetPageByID(c, int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, page)
}
