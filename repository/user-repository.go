package repository

import (
	"github.com/fazriegi/go-architecture-example/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(props *model.User) error
	FindEmail(email string) (model.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(props *model.User) error {
	return r.db.Create(props).Error
}

func (r *UserRepository) FindEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}
