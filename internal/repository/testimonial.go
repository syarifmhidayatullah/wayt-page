package repository

import (
	"github.com/project/wayt-page/internal/model"
	"gorm.io/gorm"
)

type TestimonialRepository interface {
	FindAll() ([]model.Testimonial, error)
	FindActive() ([]model.Testimonial, error)
	FindByID(id uint) (*model.Testimonial, error)
	Create(t *model.Testimonial) error
	Update(t *model.Testimonial) error
	Delete(id uint) error
}

type testimonialRepository struct{ db *gorm.DB }

func NewTestimonialRepository(db *gorm.DB) TestimonialRepository {
	return &testimonialRepository{db: db}
}

func (r *testimonialRepository) FindAll() ([]model.Testimonial, error) {
	var list []model.Testimonial
	err := r.db.Order("sort_order ASC, id ASC").Find(&list).Error
	return list, err
}

func (r *testimonialRepository) FindActive() ([]model.Testimonial, error) {
	var list []model.Testimonial
	err := r.db.Where("is_active = true").Order("sort_order ASC, id ASC").Find(&list).Error
	return list, err
}

func (r *testimonialRepository) FindByID(id uint) (*model.Testimonial, error) {
	var t model.Testimonial
	return &t, r.db.First(&t, id).Error
}

func (r *testimonialRepository) Create(t *model.Testimonial) error {
	return r.db.Create(t).Error
}

func (r *testimonialRepository) Update(t *model.Testimonial) error {
	return r.db.Save(t).Error
}

func (r *testimonialRepository) Delete(id uint) error {
	return r.db.Delete(&model.Testimonial{}, id).Error
}
