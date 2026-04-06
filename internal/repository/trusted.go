package repository

import (
	"github.com/project/wayt-page/internal/model"
	"gorm.io/gorm"
)

type TrustedRepository interface {
	ListActive() ([]model.TrustedRestaurant, error)
	ListAll() ([]model.TrustedRestaurant, error)
	Create(t *model.TrustedRestaurant) error
	Update(t *model.TrustedRestaurant) error
	Delete(id uint) error
	FindByID(id uint) (*model.TrustedRestaurant, error)
}

type trustedRepository struct{ db *gorm.DB }

func NewTrustedRepository(db *gorm.DB) TrustedRepository {
	return &trustedRepository{db: db}
}

func (r *trustedRepository) ListActive() ([]model.TrustedRestaurant, error) {
	var list []model.TrustedRestaurant
	err := r.db.Where("is_active = true").Order("sort_order asc, id asc").Find(&list).Error
	return list, err
}

func (r *trustedRepository) ListAll() ([]model.TrustedRestaurant, error) {
	var list []model.TrustedRestaurant
	err := r.db.Order("sort_order asc, id asc").Find(&list).Error
	return list, err
}

func (r *trustedRepository) Create(t *model.TrustedRestaurant) error {
	return r.db.Create(t).Error
}

func (r *trustedRepository) Update(t *model.TrustedRestaurant) error {
	return r.db.Save(t).Error
}

func (r *trustedRepository) Delete(id uint) error {
	return r.db.Delete(&model.TrustedRestaurant{}, id).Error
}

func (r *trustedRepository) FindByID(id uint) (*model.TrustedRestaurant, error) {
	var t model.TrustedRestaurant
	return &t, r.db.First(&t, id).Error
}
