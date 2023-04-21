package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/qudecim/password-manager-backend/pkg/models"
	"github.com/qudecim/password-manager-backend/pkg/repository"
)

const salt = "laadff"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int64, error) {
	user.Password = s.hash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) hash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
