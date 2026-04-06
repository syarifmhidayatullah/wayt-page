package service

import (
	"errors"

	"github.com/project/wayt-page/internal/model"
	"github.com/project/wayt-page/internal/repository"
)

type PricingService interface {
	ListAll() ([]model.PricingPlan, error)
	ListActive() ([]model.PricingPlan, error)
	GetByID(id uint) (*model.PricingPlan, error)
	Create(name string, priceMonthly, priceYearly int, features []string, isPopular bool, sortOrder int) (*model.PricingPlan, error)
	Update(id uint, name string, priceMonthly, priceYearly int, features []string, isPopular, isActive bool, sortOrder int) (*model.PricingPlan, error)
	Delete(id uint) error
}

type pricingService struct{ repo repository.PricingRepository }

func NewPricingService(repo repository.PricingRepository) PricingService {
	return &pricingService{repo: repo}
}

func (s *pricingService) ListAll() ([]model.PricingPlan, error) {
	return s.repo.FindAll()
}

func (s *pricingService) ListActive() ([]model.PricingPlan, error) {
	return s.repo.FindActive()
}

func (s *pricingService) GetByID(id uint) (*model.PricingPlan, error) {
	return s.repo.FindByID(id)
}

func (s *pricingService) Create(name string, priceMonthly, priceYearly int, features []string, isPopular bool, sortOrder int) (*model.PricingPlan, error) {
	if name == "" {
		return nil, errors.New("nama paket wajib diisi")
	}
	p := &model.PricingPlan{
		Name:         name,
		PriceMonthly: priceMonthly,
		PriceYearly:  priceYearly,
		Features:     features,
		IsPopular:    isPopular,
		IsActive:     true,
		SortOrder:    sortOrder,
	}
	return p, s.repo.Create(p)
}

func (s *pricingService) Update(id uint, name string, priceMonthly, priceYearly int, features []string, isPopular, isActive bool, sortOrder int) (*model.PricingPlan, error) {
	p, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("paket tidak ditemukan")
	}
	p.Name = name
	p.PriceMonthly = priceMonthly
	p.PriceYearly = priceYearly
	p.Features = features
	p.IsPopular = isPopular
	p.IsActive = isActive
	p.SortOrder = sortOrder
	return p, s.repo.Update(p)
}

func (s *pricingService) Delete(id uint) error {
	if _, err := s.repo.FindByID(id); err != nil {
		return errors.New("paket tidak ditemukan")
	}
	return s.repo.Delete(id)
}
