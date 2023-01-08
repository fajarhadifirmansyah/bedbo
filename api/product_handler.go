package api

import (
	"net/http"

	"github.com/fajarhadifirmansyah/bedbo/store"
	"github.com/fajarhadifirmansyah/bedbo/types"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	store store.ProductStorer
}

func NewProductHandler(pStore store.ProductStorer) *ProductHandler {
	return &ProductHandler{
		store: pStore,
	}
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	uri := URI{}
	if err := c.BindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var prod []types.Product
	if err := h.store.FindAll(&prod); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	res := types.Response{
		StatusCode: http.StatusOK,
		Data:       prod,
	}
	c.JSON(http.StatusOK, res)
}
