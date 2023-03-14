package v1

import (
	"net/http"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/yungen-lu/shared-key-value-list-system/internal/domain"
	"github.com/yungen-lu/shared-key-value-list-system/internal/usecase"
)

type pageRoutes struct {
	list usecase.List
}

func NewPageRoutes(handler *gin.RouterGroup, listUseCase usecase.List) {
	r := &pageRoutes{
		list: listUseCase,
	}
	h := handler.Group("/page")
	{
		h.GET("", r.getAll)
		h.GET("/:key", r.getByKey)
		h.POST("", r.create)
		h.PUT("/:key", r.update)
	}
}

// @Summary		Get all pages
// @Description	test
// @Accept			json
// @Produce		json
// @Success		200	{array}	domain.Page
// @Falure			500 {object}
// @Router			/page [get]
func (r *pageRoutes) getAll(c *gin.Context) {
	pages, err := r.list.GetPages(c)
	if err != nil {
		log.Error("http - v1 - page - getAll - GetPages", "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, pages)
}

// @Summary		Get a page by key
// @Description	test
// @Accept			json
// @Produce		json
// @Param			key	path		string	true	"page Key"
// @Success		200	{object}	domain.Page
// @Falure			500 {object}
// @Router			/page/{key} [get]
func (r *pageRoutes) getByKey(c *gin.Context) {
	key, ok := c.Params.Get("key")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	page, err := r.list.GetPageByKey(c, key)
	if err != nil {
		log.Error("http - v1 - page - getByKey - GetPageByKey", "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, page)
}

// @Summary		Get a page by id
// @Description	test
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"page ID"
// @Success		200	{object}	domain.Page
// @Falure			500 {object}
// @Router			/page/{id} [get]
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
		log.Error("http - v1 - page - getByID - GetPageByID", "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, page)
}

type createPageRequest struct {
	Key        string `json:"key" binding:"required"`
	NextPageID *int32 `json:"next_page_id"`
}

// @Summary		Create a page
// @Description	test
// @Accept			json
// @Produce		json
// @Param
// @Success		200	{object}	domain.Page
// @Falure			500 {object}
// @Router			/page [post]

func (r *pageRoutes) create(c *gin.Context) {
	var createPageRequest createPageRequest
	if err := c.ShouldBindJSON(&createPageRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	err := r.list.CreatePage(c, domain.Page{Key: createPageRequest.Key, NextPageID: createPageRequest.NextPageID})
	if err != nil {
		log.Error("http - v1 - page - create - CreatePage", "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

type updatePageRequest struct {
	Key        string `json:"key"`
	NextPageID *int32 `json:"next_page_id"`
}

// @Summary		Update a page
// @Description	test
// @Accept			json
// @Produce		json
// @Param page body updatePageRequest true "Page"
// @Success		200	{object}	domain.Page
// @Falure			500 {object}
// @Router			/page{key} [put]

func (r *pageRoutes) update(c *gin.Context) {
	key, ok := c.Params.Get("key")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	var updatePageRequest updatePageRequest
	if err := c.ShouldBindJSON(&updatePageRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	err := r.list.UpdatePageByKey(c, key, domain.Page{Key: updatePageRequest.Key, NextPageID: updatePageRequest.NextPageID})
	if err != nil {
		log.Error("http - v1 - page - update - UpdatePageByKey", "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
