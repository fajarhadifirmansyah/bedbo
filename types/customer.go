package types

import (
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tabler interface {
	TableName() string
}

type Customer struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name      string    `json:"name" gorm:"column:name"`
	Address   string    `json:"address" gorm:"column:address"`
	HandPhone string    `json:"handphone" gorm:"column:no_handphone" `
	Gender    string    `json:"gender" gorm:"column:gender" `
}

func (Customer) TableName() string {
	return "mst_customer"
}

func (c *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}

type CustomerDto struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	HandPhone string    `json:"handphone"`
	Gender    string    `json:"gender"`
}

type CustomerReq struct {
	Name      string `json:"name" binding:"required"`
	Address   string `json:"address"`
	HandPhone string `json:"handphone" binding:"required"`
	Gender    string `json:"gender" binding:"required,gender"`
}

func ConvertToCustDTO(c *Customer) (*CustomerDto, error) {
	return &CustomerDto{
		ID:        c.ID,
		Name:      c.Name,
		Address:   c.Address,
		HandPhone: c.HandPhone,
		Gender:    c.Gender,
	}, nil
}

func ConvertCustReqToCustEntity(c *CustomerReq) (*Customer, error) {
	return &Customer{
		Name:      c.Name,
		Address:   c.Address,
		HandPhone: c.HandPhone,
		Gender:    strings.ToUpper(c.Gender),
	}, nil
}
