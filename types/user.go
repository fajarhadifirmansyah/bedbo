package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Email    string    `json:"email" gorm:"column:email"`
	Password string    `json:"-" gorm:"column:password"`
	Role     string    `json:"role" gorm:"column:role" `
}

func (User) TableName() string {
	return "user"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

type SignUpRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

func ConvertSignUpReqToCustEntity(p *SignUpRequest) (*User, error) {
	return &User{
		Email:    p.Email,
		Password: p.Password,
		Role:     p.Role,
	}, nil
}

type SignInRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
