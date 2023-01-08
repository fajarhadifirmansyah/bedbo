package store

import (
	"fmt"

	"github.com/fajarhadifirmansyah/bedbo/log"
	"github.com/fajarhadifirmansyah/bedbo/types"
	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (s *UserStore) FindUserById(id string, u *types.User) error {
	l := log.Get()
	if err := s.db.First(&u, "id = ?", id).Error; err != nil {
		l.Error().Err(err).Msg("something wrong on database")
		return fmt.Errorf(err.Error())
	}

	return nil
}

func (s *UserStore) FindUserByEmail(email string, u *types.User) error {
	l := log.Get()
	if err := s.db.First(&u, "email = ?", email).Error; err != nil {
		l.Error().Err(err).Msg("something wrong on database")
		return fmt.Errorf(err.Error())
	}
	return nil
}
