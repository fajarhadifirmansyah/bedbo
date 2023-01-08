package types

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Order struct {
	// gorm.Model
	ID           int64           `gorm:"primary_key;type:int;autoIncrement;not_null" json:"id"`
	OrderDate    time.Time       `json:"order_date" gorm:"column:order_date"`
	Total        decimal.Decimal `json:"total" gorm:"column:total"`
	Status       string          `json:"status" gorm:"column:status" `
	CustomerID   uuid.UUID       `json:"-" gorm:"column:customer_id" `
	Customer     Customer        `json:"customer"`
	OrderDetails []OrderDetail   `json:"details" gorm:"foreignKey:OrderID;references:ID"`
}

func (Order) TableName() string {
	return "trx_order"
}

type OrderDetail struct {
	ID        uuid.UUID       `json:"id" gorm:"primary_key;not_null"`
	UnitPrice decimal.Decimal `json:"unit_price" gorm:"column:unit_price"`
	Qty       int64           `json:"qty" gorm:"column:qty"`
	Total     decimal.Decimal `json:"total" gorm:"column:total" `
	ProductID uuid.UUID       `json:"-" gorm:"column:product_id" `
	Product   Product         `json:"product" gorm:"foreignKey:ProductID;references:ID"`
	OrderID   int64           `json:"-" gorm:"column:order_id" `
	// Order   Order `json:"order"`
}

func (OrderDetail) TableName() string {
	return "trx_order_detail"
}

func (c *OrderDetail) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}

type OrderReq struct {
	OrderDate    time.Time        `json:"order_date"`
	Total        decimal.Decimal  `json:"total" binding:"required"`
	Status       string           `json:"status" binding:"required"`
	CustomerID   uuid.UUID        `json:"customer_id" binding:"required"`
	OrderDetails []OrderDetailReq `json:"details"`
}

type OrderDetailReq struct {
	UnitPrice decimal.Decimal `json:"unit_price" binding:"required"`
	Qty       int64           `json:"qty" binding:"required"`
	Total     decimal.Decimal `json:"total" binding:"required"`
	ProductID uuid.UUID       `json:"product_id" binding:"required"`
}

type UpdateStatusReq struct {
	Status string `json:"status" binding:"required"`
}
