package repository

import (
	"github.com/project/wayt-page/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsername(username string) (*model.User, error)
	FindByID(id uint) (*model.User, error)
	Create(u *model.User) error
	Update(u *model.User) error
	Delete(id uint) error
	ListAll() ([]model.User, error)
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

func (r *userRepository) FindByID(id uint) (*model.User, error) {
	var u model.User
	return &u, r.db.First(&u, id).Error
}

func (r *userRepository) Update(u *model.User) error {
	return r.db.Save(u).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}

func (r *userRepository) ListAll() ([]model.User, error) {
	var list []model.User
	err := r.db.Order("id asc").Find(&list).Error
	return list, err
}

func (r *userRepository) ExistsAny() (bool, error) {
	var count int64
	err := r.db.Model(&model.User{}).Count(&count).Error
	return count > 0, err
}
