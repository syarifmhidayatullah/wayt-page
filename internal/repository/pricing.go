package repository

import (
	"github.com/project/wayt-page/internal/model"
	"gorm.io/gorm"
)

type PricingRepository interface {
	FindAll() ([]model.PricingPlan, error)
	FindActive() ([]model.PricingPlan, error)
	FindByID(id uint) (*model.PricingPlan, error)
	Create(p *model.PricingPlan) error
	Update(p *model.PricingPlan) error
	Delete(id uint) error
}

type pricingRepository struct{ db *gorm.DB }

func NewPricingRepository(db *gorm.DB) PricingRepository {
	return &pricingRepository{db: db}
}

func (r *pricingRepository) FindAll() ([]model.PricingPlan, error) {
	var list []model.PricingPlan
	err := r.db.Order("sort_order ASC").Find(&list).Error
	return list, err
}

func (r *pricingRepository) FindActive() ([]model.PricingPlan, error) {
	var list []model.PricingPlan
	err := r.db.Where("is_active = true").Order("sort_order ASC").Find(&list).Error
	return list, err
}

func (r *pricingRepository) FindByID(id uint) (*model.PricingPlan, error) {
	var p model.PricingPlan
	return &p, r.db.First(&p, id).Error
}

func (r *pricingRepository) Create(p *model.PricingPlan) error {
	return r.db.Create(p).Error
}

func (r *pricingRepository) Update(p *model.PricingPlan) error {
	return r.db.Save(p).Error
}

func (r *pricingRepository) Delete(id uint) error {
	return r.db.Delete(&model.PricingPlan{}, id).Error
}
