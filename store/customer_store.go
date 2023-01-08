package store

import (
	"fmt"

	"github.com/fajarhadifirmansyah/bedbo/log"
	"github.com/fajarhadifirmansyah/bedbo/types"
	"gorm.io/gorm"
)

type CustomerStore struct {
	db *gorm.DB
}

func NewCustomerStore(db *gorm.DB) *CustomerStore {
	return &CustomerStore{
		db: db,
	}
}

func (s *CustomerStore) FindByID(id string, c *types.Customer) error {
	l := log.Get()
	if err := s.db.First(&c, "id = ?", id).Error; err != nil {
		l.Error().Err(err).Msg("something wrong on database")
		return fmt.Errorf(err.Error())
	}

	return nil
}

func (s *CustomerStore) Insert(c *types.Customer) error {
	l := log.Get()
	if err := s.db.Create(c).Error; err != nil {
		l.Error().Err(err).Msg("something wrong on database")
		return fmt.Errorf(err.Error())
	}
	return nil
}

func (s *CustomerStore) Delete(id string) error {
	var cExist types.Customer
	if err := s.FindByID(id, &cExist); err != nil {
		return err
	}
	s.db.Where("id = ?", cExist.ID).Delete(&types.Customer{})
	return nil
}

func (s *CustomerStore) Update(c *types.Customer) error {
	s.db.Save(c)
	return nil
}

func (s *CustomerStore) PaginateOffsetLimit(reqPaging types.PagingReqBasic) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pageSize, order, offset := reqPaging.SetPagingParam(types.Customer{}, "name", "asc")
		return db.Offset(*offset).Limit(*pageSize).Order(order)
	}
}

func (s *CustomerStore) PaginateSearchAndFilter(search string, filter interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		f := filter.(map[string]interface{})
		var fGender []string
		if f["gender"] != nil {
			fGender = f["gender"].([]string)
		}

		search = "%" + search + "%"
		chain := db.Where("name LIKE ? OR address LIKE ? OR no_handphone LIKE ?", search, search, search)
		if len(fGender) > 0 {
			chain = chain.Where("gender IN ?", fGender)
		}

		return chain
	}
}

func (s *CustomerStore) Paginate(reqPaging types.PagingReqBasic,
	data interface{}, count *int64, filter interface{}) {
	s.db.Scopes(s.PaginateOffsetLimit(reqPaging), s.PaginateSearchAndFilter(reqPaging.Search, filter)).Find(data)
	s.db.Model(&types.Customer{}).Scopes(s.PaginateOffsetLimit(reqPaging), s.PaginateSearchAndFilter(reqPaging.Search, filter)).Count(count)
}
