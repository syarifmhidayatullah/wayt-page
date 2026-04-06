package repository

import (
	"github.com/project/wayt-page/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsername(username string) (*model.User, error)
	Create(u *model.User) error
	ExistsAny() (bool, error)
}

type userRepository struct{ db *gorm.DB }

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByUsername(username string) (*model.User, error) {
	var u model.User
	return &u, r.db.Where("username = ?", username).First(&u).Error
}

func (r *userRepository) Create(u *model.User) error {
	return r.db.Create(u).Error
}

func (r *userRepository) ExistsAny() (bool, error) {
	var count int64
	err := r.db.Model(&model.User{}).Count(&count).Error
	return count > 0, err
}
