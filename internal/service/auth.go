package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/project/wayt-page/internal/model"
	"github.com/project/wayt-page/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(username, password string) (string, error)
	SeedAdmin(username, password string) error
	ListAdmins() ([]model.User, error)
	CreateAdmin(username, password string) (*model.User, error)
	UpdatePassword(id uint, newPassword string) error
	DeleteAdmin(id uint) error
}

type authService struct {
	repo      repository.UserRepository
	jwtSecret string
}

func NewAuthService(repo repository.UserRepository, jwtSecret string) AuthService {
	return &authService{repo: repo, jwtSecret: jwtSecret}
}

func (s *authService) Login(username, password string) (string, error) {
	u, err := s.repo.FindByUsername(username)
	if err != nil {
		return "", errors.New("username atau password salah")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return "", errors.New("username atau password salah")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      u.ID,
		"username": u.Username,
		"exp":      time.Now().Add(8 * time.Hour).Unix(),
	})
	return token.SignedString([]byte(s.jwtSecret))
}

func (s *authService) SeedAdmin(username, password string) error {
	exists, err := s.repo.ExistsAny()
	if err != nil || exists {
		return err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return s.repo.Create(&model.User{Username: username, Password: string(hash)})
}

func (s *authService) ListAdmins() ([]model.User, error) {
	return s.repo.ListAll()
}

func (s *authService) CreateAdmin(username, password string) (*model.User, error) {
	if _, err := s.repo.FindByUsername(username); err == nil {
		return nil, errors.New("username sudah digunakan")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u := &model.User{Username: username, Password: string(hash)}
	return u, s.repo.Create(u)
}

func (s *authService) UpdatePassword(id uint, newPassword string) error {
	u, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("admin tidak ditemukan")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return s.repo.Update(u)
}

func (s *authService) DeleteAdmin(id uint) error {
	count, err := s.repo.ListAll()
	if err != nil {
		return err
	}
	if len(count) <= 1 {
		return errors.New("tidak bisa menghapus admin terakhir")
	}
	return s.repo.Delete(id)
}
