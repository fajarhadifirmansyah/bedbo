package store

import (
	"fmt"

	"github.com/fajarhadifirmansyah/bedbo/log"
	"github.com/fajarhadifirmansyah/bedbo/types"
	"gorm.io/gorm"
)

type OrderStore struct {
	db *gorm.DB
}

func NewOrderStore(db *gorm.DB) *OrderStore {
	return &OrderStore{
		db: db,
	}
}

func (s *OrderStore) PaginateOffsetLimit(reqPaging types.PagingReqBasic) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pageSize, order, offset := reqPaging.SetPagingParam(types.Order{}, "order_date", "desc")
		return db.Offset(*offset).Limit(*pageSize).Order(order)
	}
}

func (s *OrderStore) PaginateSearchAndFilter(search string, filter interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		f := filter.(map[string]interface{})
		var fStatus []string
		if f["status"] != nil {
			fStatus = f["status"].([]string)
		}

		search = "%" + search + "%"
		// chain := db.Joins("JOIN mst_customer on c.id=order.customer_id").Where("mst_customer.name LIKE ?", search)
		chain := db.Where("trx_order.status LIKE ? OR \"Customer\".\"name\" LIKE ?", search, search)
		if len(fStatus) > 0 {
			chain = chain.Where("status IN ?", fStatus)
		}

		return chain
	}
}

func (s *OrderStore) Paginate(reqPaging types.PagingReqBasic,
	data interface{}, count *int64, filter interface{}) {
	s.db.Joins("Customer").Scopes(s.PaginateOffsetLimit(reqPaging), s.PaginateSearchAndFilter(reqPaging.Search, filter)).Find(data)
	s.db.Model(&types.Order{}).Joins("Customer").Scopes(s.PaginateOffsetLimit(reqPaging), s.PaginateSearchAndFilter(reqPaging.Search, filter)).Count(count)
}

func (s *OrderStore) FindByID(id int64, o *types.Order) error {
	l := log.Get()
	if err := s.db.Preload("Customer").Preload("OrderDetails").Preload("OrderDetails.Product").First(&o, id).Error; err != nil {
		l.Error().Err(err).Msg("something wrong on database")
		return fmt.Errorf(err.Error())
	}

	return nil
}

func (s *OrderStore) InsertOrder(o *types.Order) error {
	l := log.Get()
	if err := s.db.Create(o).Error; err != nil {
		l.Error().Err(err).Msg("something wrong on database")
		return fmt.Errorf(err.Error())
	}
	return nil
}

func (s *OrderStore) InsertOrderDetail(o *types.OrderDetail) error {
	l := log.Get()
	if err := s.db.Create(o).Error; err != nil {
		l.Error().Err(err).Msg("something wrong on database")
		return fmt.Errorf(err.Error())
	}
	return nil
}

func (s *OrderStore) UpdateStatusOrder(id int64, status string) error {

	var oExist types.Order
	if err := s.FindByID(id, &oExist); err != nil {
		return err
	}

	oExist.Status = status
	s.db.Save(&oExist)
	return nil
}

func (s *OrderStore) Delete(id int64) error {
	var oExist types.Order
	if err := s.FindByID(id, &oExist); err != nil {
		return err
	}
	s.db.Where("id = ?", oExist.ID).Delete(&types.Order{})
	return nil
}
