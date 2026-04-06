package repository

import (
	"github.com/project/wayt-page/internal/model"
	"gorm.io/gorm"
)

type SettingRepository interface {
	FindAll() ([]model.Setting, error)
	FindByKey(key string) (*model.Setting, error)
	Upsert(key, value string) error
}

type settingRepository struct{ db *gorm.DB }

func NewSettingRepository(db *gorm.DB) SettingRepository {
	return &settingRepository{db: db}
}

func (r *settingRepository) FindAll() ([]model.Setting, error) {
	var list []model.Setting
	err := r.db.Order("key ASC").Find(&list).Error
	return list, err
}

func (r *settingRepository) FindByKey(key string) (*model.Setting, error) {
	var s model.Setting
	return &s, r.db.Where("key = ?", key).First(&s).Error
}

func (r *settingRepository) Upsert(key, value string) error {
	return r.db.Model(&model.Setting{}).Where("key = ?", key).
		Update("value", value).Error
}
