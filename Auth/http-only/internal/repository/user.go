package repository

import (
	"github.com/Rai-Sahil/backend/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (domain.User, error)
	Create(user domain.User) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{db}
}

func (r *userRepo) FindByEmail(email string) (domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *userRepo) Create(user domain.User) error {
	return r.db.Create(&user).Error
}
