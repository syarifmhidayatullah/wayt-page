package service

import (
	"errors"
	"strings"

	"github.com/project/wayt-page/internal/model"
	"github.com/project/wayt-page/internal/repository"
)

type TrustedService interface {
	ListActive() ([]model.TrustedRestaurant, error)
	ListAll() ([]model.TrustedRestaurant, error)
	Create(name, emoji, rating string, sortOrder int) (*model.TrustedRestaurant, error)
	Update(id uint, name, emoji, rating string, isActive bool, sortOrder int) (*model.TrustedRestaurant, error)
	Delete(id uint) error
}

type trustedService struct{ repo repository.TrustedRepository }

func NewTrustedService(repo repository.TrustedRepository) TrustedService {
	return &trustedService{repo: repo}
}

func (s *trustedService) ListActive() ([]model.TrustedRestaurant, error) {
	return s.repo.ListActive()
}

func (s *trustedService) ListAll() ([]model.TrustedRestaurant, error) {
	return s.repo.ListAll()
}

func (s *trustedService) Create(name, emoji, rating string, sortOrder int) (*model.TrustedRestaurant, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("nama restoran wajib diisi")
	}
	if emoji == "" {
		emoji = "🍽️"
	}
	if rating == "" {
		rating = "5.0"
	}
	t := &model.TrustedRestaurant{
		Name:      strings.TrimSpace(name),
		Emoji:     emoji,
		Rating:    rating,
		IsActive:  true,
		SortOrder: sortOrder,
	}
	return t, s.repo.Create(t)
}

func (s *trustedService) Update(id uint, name, emoji, rating string, isActive bool, sortOrder int) (*model.TrustedRestaurant, error) {
	t, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("restoran tidak ditemukan")
	}
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("nama restoran wajib diisi")
	}
	t.Name = strings.TrimSpace(name)
	t.Emoji = emoji
	t.Rating = rating
	t.IsActive = isActive
	t.SortOrder = sortOrder
	return t, s.repo.Update(t)
}

func (s *trustedService) Delete(id uint) error {
	return s.repo.Delete(id)
}
