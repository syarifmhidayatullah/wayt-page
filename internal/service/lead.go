package service

import (
	"errors"
	"strings"

	"github.com/project/wayt-page/internal/model"
	"github.com/project/wayt-page/internal/repository"
)

type LeadService interface {
	Submit(fullName, restaurantName, email, phone string) (*model.Lead, error)
	ListAll() ([]model.Lead, error)
	Delete(id uint) error
}

type leadService struct{ repo repository.LeadRepository }

func NewLeadService(repo repository.LeadRepository) LeadService {
	return &leadService{repo: repo}
}

func (s *leadService) Submit(fullName, restaurantName, email, phone string) (*model.Lead, error) {
	if strings.TrimSpace(fullName) == "" || strings.TrimSpace(restaurantName) == "" || strings.TrimSpace(email) == "" {
		return nil, errors.New("nama, restoran, dan email wajib diisi")
	}
	l := &model.Lead{
		FullName:       strings.TrimSpace(fullName),
		RestaurantName: strings.TrimSpace(restaurantName),
		Email:          strings.TrimSpace(email),
		Phone:          strings.TrimSpace(phone),
	}
	return l, s.repo.Create(l)
}

func (s *leadService) ListAll() ([]model.Lead, error) {
	return s.repo.ListAll()
}

func (s *leadService) Delete(id uint) error {
	return s.repo.Delete(id)
}
