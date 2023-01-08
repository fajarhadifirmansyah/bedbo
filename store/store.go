package store

import (
	"github.com/fajarhadifirmansyah/bedbo/types"
	"gorm.io/gorm"
)

type Paging interface {
	PaginateOffsetLimit(reqPaging types.PagingReqBasic) func(db *gorm.DB) *gorm.DB
	PaginateSearchAndFilter(search string, filter interface{}) func(db *gorm.DB) *gorm.DB
}

type PaginationStorer interface {
	Paging
	Paginate(reqPaging types.PagingReqBasic,
		data interface{}, count *int64, filter interface{})
}

type CustomerStorer interface {
	PaginationStorer
	Delete(id string) error
	FindByID(id string, c *types.Customer) error
	Insert(c *types.Customer) error
	Update(c *types.Customer) error
}

type ProductStorer interface {
	FindAll(p *[]types.Product) error
}

type OrderStorer interface {
	PaginationStorer
	FindByID(id int64, o *types.Order) error
	InsertOrder(o *types.Order) error
	InsertOrderDetail(o *types.OrderDetail) error
	UpdateStatusOrder(id int64, status string) error
	Delete(id int64) error
}

type UserStorer interface {
	FindUserById(id string, u *types.User) error
	FindUserByEmail(email string, u *types.User) error
}

type Authstorer interface {
	SignUpUser(u *types.User) error
	SignInUser(req *types.SignInRequest, u *types.User) error
}
