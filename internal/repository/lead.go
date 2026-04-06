package repository

import (
	"github.com/project/wayt-page/internal/model"
	"gorm.io/gorm"
)

type LeadRepository interface {
	Create(l *model.Lead) error
	ListAll() ([]model.Lead, error)
	Delete(id uint) error
}

type leadRepository struct{ db *gorm.DB }

func NewLeadRepository(db *gorm.DB) LeadRepository {
	return &leadRepository{db: db}
}

func (r *leadRepository) Create(l *model.Lead) error {
	return r.db.Create(l).Error
}

func (r *leadRepository) ListAll() ([]model.Lead, error) {
	var list []model.Lead
	err := r.db.Order("created_at desc").Find(&list).Error
	return list, err
}

func (r *leadRepository) Delete(id uint) error {
	return r.db.Delete(&model.Lead{}, id).Error
}
