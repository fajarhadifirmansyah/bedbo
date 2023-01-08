package types

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	ID       uuid.UUID       `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name     string          `json:"name" gorm:"column:name"`
	Desc     string          `json:"description" gorm:"column:description"`
	Unit     string          `json:"unit" gorm:"column:unit" `
	Price    decimal.Decimal `json:"price" gorm:"column:price" `
	Qty      int64           `json:"qty" gorm:"column:quantity" `
	Category string          `json:"category" gorm:"column:category" `
}

func (Product) TableName() string {
	return "mst_product"
}

func (c *Product) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}
