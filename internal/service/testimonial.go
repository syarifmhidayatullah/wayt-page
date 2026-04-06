package service

import (
	"errors"

	"github.com/project/wayt-page/internal/model"
	"github.com/project/wayt-page/internal/repository"
)

type TestimonialService interface {
	ListAll() ([]model.Testimonial, error)
	ListActive() ([]model.Testimonial, error)
	Create(name, restaurant, quote, phone string, rating, sortOrder int) (*model.Testimonial, error)
	Update(id uint, name, restaurant, quote, phone string, rating, sortOrder int, isActive bool) (*model.Testimonial, error)
	Delete(id uint) error
}

type testimonialService struct{ repo repository.TestimonialRepository }

func NewTestimonialService(repo repository.TestimonialRepository) TestimonialService {
	return &testimonialService{repo: repo}
}

func (s *testimonialService) ListAll() ([]model.Testimonial, error) {
	return s.repo.FindAll()
}

func (s *testimonialService) ListActive() ([]model.Testimonial, error) {
	return s.repo.FindActive()
}

func (s *testimonialService) Create(name, restaurant, quote, phone string, rating, sortOrder int) (*model.Testimonial, error) {
	if name == "" || quote == "" {
		return nil, errors.New("nama dan kutipan wajib diisi")
	}
	t := &model.Testimonial{
		Name: name, Restaurant: restaurant, Quote: quote,
		Phone: phone, Rating: rating, SortOrder: sortOrder, IsActive: true,
	}
	return t, s.repo.Create(t)
}

func (s *testimonialService) Update(id uint, name, restaurant, quote, phone string, rating, sortOrder int, isActive bool) (*model.Testimonial, error) {
	t, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("testimoni tidak ditemukan")
	}
	t.Name = name
	t.Restaurant = restaurant
	t.Quote = quote
	t.Phone = phone
	t.Rating = rating
	t.SortOrder = sortOrder
	t.IsActive = isActive
	return t, s.repo.Update(t)
}

func (s *testimonialService) Delete(id uint) error {
	if _, err := s.repo.FindByID(id); err != nil {
		return errors.New("testimoni tidak ditemukan")
	}
	return s.repo.Delete(id)
}
