package store

import (
	"fmt"
	"strings"

	"github.com/fajarhadifirmansyah/bedbo/log"
	"github.com/fajarhadifirmansyah/bedbo/types"
	"gorm.io/gorm"
)

type AuthStore struct {
	db *gorm.DB
}

func NewAuthStore(db *gorm.DB) *AuthStore {
	return &AuthStore{
		db: db,
	}
}

func (s AuthStore) SignUpUser(u *types.User) error {
	l := log.Get()
	if err := s.db.Create(u).Error; err != nil {
		l.Error().Err(err).Msg("something wrong on database")

		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return fmt.Errorf("username or email is used")
		}
		return fmt.Errorf(err.Error())
	}
	return nil
}

func (s *AuthStore) SignInUser(req *types.SignInRequest, u *types.User) error {

	return nil
}
