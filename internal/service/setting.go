package service

import (
	"github.com/project/wayt-page/internal/model"
	"github.com/project/wayt-page/internal/repository"
)

type SettingService interface {
	ListAll() ([]model.Setting, error)
	GetMap() (map[string]string, error)
	Set(key, value string) error
}

type settingService struct{ repo repository.SettingRepository }

func NewSettingService(repo repository.SettingRepository) SettingService {
	return &settingService{repo: repo}
}

func (s *settingService) ListAll() ([]model.Setting, error) {
	return s.repo.FindAll()
}

func (s *settingService) GetMap() (map[string]string, error) {
	list, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	m := make(map[string]string, len(list))
	for _, s := range list {
		m[s.Key] = s.Value
	}
	return m, nil
}

func (s *settingService) Set(key, value string) error {
	return s.repo.Upsert(key, value)
}
