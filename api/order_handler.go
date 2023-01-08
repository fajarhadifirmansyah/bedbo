package api

import (
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/fajarhadifirmansyah/bedbo/store"
	"github.com/fajarhadifirmansyah/bedbo/types"
	"github.com/fajarhadifirmansyah/bedbo/utils"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	store store.OrderStorer
}

func NewOrderHandler(oStore store.OrderStorer) *OrderHandler {
	return &OrderHandler{
		store: oStore,
	}
}

func (h *OrderHandler) GetOrderHandler(c *gin.Context) {
	// c.JSON(http.StatusOK, "res")
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("limit"))
	column := c.Query("orderby")
	sort := c.Query("sort")
	search := c.Query("search")

	reqPaging := types.PagingReqBasic{
		Page:     &page,
		PageSize: &pageSize,
		Column:   column,
		Sort:     sort,
		Search:   search,
	}

	status := c.Query("status")
	filter := make(map[string]interface{})
	if status != "" {
		filter["status"] = strings.Split(status, ",")
	}

	var d []types.Order
	var count int64
	h.store.Paginate(reqPaging, &d, &count, filter)

	res := types.ResponsePaging{
		Response: types.Response{
			StatusCode: http.StatusOK,
			Data:       d,
		},
		TotalItems:  count,
		TotalPage:   int64(math.Ceil(float64(count) / float64(pageSize))),
		PageSize:    int64(pageSize),
		CurrentPage: int64(page),
	}
	c.JSON(http.StatusOK, res)
}

func (h *OrderHandler) GetByIDHandler(c *gin.Context) {
	id := c.Param("orderID")
	pageParam, _ := strconv.ParseInt(id, 10, 64)

	var order types.Order
	if err := h.store.FindByID(pageParam, &order); err != nil {
		if err.Error() == "record not found" || strings.Contains(err.Error(), "invalid input syntax for type uuid") {
			c.JSON(http.StatusNotFound, gin.H{
				"data":    nil,
				"message": "data not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) CreateOrderHandler(c *gin.Context) {

	var req types.OrderReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ParseError(err)})
		return
	}

	order := types.Order{
		OrderDate:  time.Now(),
		Total:      req.Total,
		Status:     req.Status,
		CustomerID: req.CustomerID,
	}

	if err := h.store.InsertOrder(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(req.OrderDetails) > 0 {

		for _, odReq := range req.OrderDetails {
			od := types.OrderDetail{
				UnitPrice: odReq.UnitPrice,
				Qty:       odReq.Qty,
				Total:     odReq.Total,
				ProductID: odReq.ProductID,
				OrderID:   order.ID,
			}
			if err := h.store.InsertOrderDetail(&od); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}
	}

	c.JSON(http.StatusCreated, gin.H{"message": "created!"})
}

func (h *OrderHandler) UpdateOrderHandler(c *gin.Context) {
	idParam := c.Param("orderID")
	Id, _ := strconv.ParseInt(idParam, 10, 64)

	var req types.UpdateStatusReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ParseError(err)})
		return
	}

	if err := h.store.UpdateStatusOrder(Id, req.Status); err != nil {
		if err.Error() == "record not found" || strings.Contains(err.Error(), "invalid input syntax for type uuid") {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad Request",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated " + idParam})
}

func (h *OrderHandler) DeleteOrderHandler(c *gin.Context) {
	idParam := c.Param("orderID")
	Id, _ := strconv.ParseInt(idParam, 10, 64)

	if err := h.store.Delete(Id); err != nil {
		if err.Error() == "record not found" || strings.Contains(err.Error(), "invalid input syntax for type uuid") {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad Request",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted " + idParam})
}
