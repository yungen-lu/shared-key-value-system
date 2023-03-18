package v1

import (
	"net/http"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/yungen-lu/shared-key-value-list-system/internal/domain"
	"github.com/yungen-lu/shared-key-value-list-system/internal/usecase"
)

type headRoutes struct {
	list usecase.List
}

func NewHeadRoutes(handler *gin.RouterGroup, listUseCase usecase.List) {
	r := &headRoutes{
		list: listUseCase,
	}
	h := handler.Group("/head")
	{
		h.GET("", r.getAll)
		h.GET("/:key", r.getByKey)
		h.POST("", r.create)
		h.PUT("/:key", r.update)
		h.DELETE("/:key", r.delete)
	}
}

// @Summary		Get all heads
// @Description	test
// @Accept			json
// @Produce		json
// @Success		200	{array}	domain.List
// @Falure			500 {object}
// @Router			/head [get]
func (r *headRoutes) getAll(c *gin.Context) {
	lists, err := r.list.GetHeads(c)
	if err != nil {
		log.Error("http - v1 - head - getAll - GetHeads", err)
		// c.JSON(http.StatusInternalServerError, gin.H{})
		responseError(c, err)
		return
	}
	c.JSON(http.StatusOK, lists)
}

// @Summary		Get a head by key
// @Description	test
// @Accept			json
// @Produce		json
// @Param			key	path		int	true	"Head Key"
// @Success		200	{object}	domain.List
// @Falure			500 {object}
// @Router			/head/{key} [get]
func (r *headRoutes) getByKey(c *gin.Context) {
	key, ok := c.Params.Get("key")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	list, err := r.list.GetHeadByKey(c, key)
	if err != nil {
		log.Error("http - v1 - head - getByKey - GetHeadByKey", "err", err, "key", key)
		// c.JSON(http.StatusInternalServerError, gin.H{})
		responseError(c, err)
		return
	}
	c.JSON(http.StatusOK, list)
}

// @Summary		Get a head by id
// @Description	test
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Head ID"
// @Success		200	{object}	domain.List
// @Falure			500 {object}
// @Router			/head/{id} [get]
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
		log.Error("http - v1 - head - getByID - GetHeadByID", "err", err)
		// c.JSON(http.StatusInternalServerError, gin.H{})
		responseError(c, err)
		return
	}
	c.JSON(http.StatusOK, list)
}

type CreateHeadRequest struct {
	// ID         int32 `json:"id" binding:"required"`
	Key         string  `json:"key" binding:"required"`
	NextPageKey *string `json:"next_page_key"`
}

// @Summary		Create a head
// @Description	test
// @Accept			json
// @Produce		json
// @Param
// @Success		200	{object}	domain.List
// @Falure			500 {object}
// @Router			/head [post]
func (r *headRoutes) create(c *gin.Context) {
	var req CreateHeadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	err := r.list.CreateHead(c, domain.List{Key: req.Key, NextPageKey: req.NextPageKey})
	if err != nil {
		log.Error("http - v1 - head - create - CreateHead", "err", err)
		// c.JSON(http.StatusInternalServerError, gin.H{})
		responseError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

type UpdateHeadRequest struct {
	NextPageKey *string `json:"next_page_key"`
}

// @Summary Update a head
// @Description test
// @Accept json
// @Produce json
// @Param head body updateHeadRequest true "Head"
// @Success 200 {object} model.Head
// @Failure 500 {object}
// @Router /head [put]

func (r *headRoutes) update(c *gin.Context) {
	key, ok := c.Params.Get("key")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	var req UpdateHeadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	err := r.list.UpdateHeadByKey(c, key, domain.List{NextPageKey: req.NextPageKey})
	if err != nil {
		log.Error("http - v1 - head - update - UpdateHeadByKey", "err", err)
		// c.JSON(http.StatusInternalServerError, gin.H{})
		responseError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// @Summary Delete a head
// @Description test
// @Success 200
// @Failure 500 {object}
func (r *headRoutes) delete(c *gin.Context) {
	key, ok := c.Params.Get("key")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	err := r.list.DeleteHeadByKey(c, key)
	if err != nil {
		log.Error("http - v1 - head - delete - DeleteHeadByKey", "err", err)
		// c.JSON(http.StatusInternalServerError, gin.H{})
		responseError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
