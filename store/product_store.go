package store

import (
	"fmt"

	"github.com/fajarhadifirmansyah/bedbo/log"
	"github.com/fajarhadifirmansyah/bedbo/types"
	"gorm.io/gorm"
)

type ProductStore struct {
	db *gorm.DB
}

func NewProductStore(db *gorm.DB) *ProductStore {
	return &ProductStore{
		db: db,
	}
}

func (s *ProductStore) FindAll(p *[]types.Product) error {
	l := log.Get()
	if err := s.db.Find(p).Error; err != nil {
		l.Error().Err(err).Msg("something wrong on database")
		return fmt.Errorf(err.Error())
	}

	return nil
}
