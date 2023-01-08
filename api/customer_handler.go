package api

import (
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/fajarhadifirmansyah/bedbo/store"
	"github.com/fajarhadifirmansyah/bedbo/types"
	"github.com/fajarhadifirmansyah/bedbo/utils"
	"github.com/gin-gonic/gin"
)

type URI struct {
	CustID string `json:"customerid" uri:"customerid"`
}

type CustomerHandler struct {
	store store.CustomerStorer
}

func NewCustomerHandler(cStore store.CustomerStorer) *CustomerHandler {
	return &CustomerHandler{
		store: cStore,
	}
}

func (h *CustomerHandler) GetCustomersHandler(c *gin.Context) {
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

	gender := c.Query("gender")
	filter := make(map[string]interface{})
	if gender != "" {
		filter["gender"] = strings.Split(gender, ",")
	}

	var d []types.Customer
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

func (h *CustomerHandler) GetCustomerByIDHandler(c *gin.Context) {
	uri := URI{}
	if err := c.BindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var cust types.Customer
	if err := h.store.FindByID(uri.CustID, &cust); err != nil {
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

	custDto, _ := types.ConvertToCustDTO(&cust)
	c.JSON(http.StatusOK, custDto)
}

func (h *CustomerHandler) CreateCustomerHandler(c *gin.Context) {

	var cr types.CustomerReq
	if err := c.ShouldBindJSON(&cr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ParseError(err)})
		return
	}

	cust, _ := types.ConvertCustReqToCustEntity(&cr)

	if err := h.store.Insert(cust); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "created!"})
}

func (h *CustomerHandler) DeleteCustomerHandler(c *gin.Context) {

	uri := URI{}
	if err := c.BindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := h.store.Delete(uri.CustID); err != nil {
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

	c.JSON(http.StatusOK, gin.H{"message": "deleted " + uri.CustID})
}

func (h *CustomerHandler) UpdateCustomerHandler(c *gin.Context) {

	id := c.Param("customerid")

	var cr types.CustomerReq
	if err := c.ShouldBindJSON(&cr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var cExist types.Customer
	if err := h.store.FindByID(id, &cExist); err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	cExist.Name = cr.Name
	cExist.Address = cr.Address
	cExist.HandPhone = cr.HandPhone
	cExist.Gender = strings.ToUpper(cr.Gender)
	h.store.Update(&cExist)

	c.JSON(http.StatusOK, gin.H{"message": "Updated " + id})
}
