package service

import (
	"github.com/qudecim/password-manager-backend/pkg/models"
	"github.com/qudecim/password-manager-backend/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int64, error)
	GenerateToken(user models.User) (string, error)
	ParseToken(token string) (int, error)
}

type Secret interface {
	GetAllSecrets(user_id int) ([]models.Secret, error)
	GetOneSecret(user_id int, secret_id int) (models.Secret, error)
	CreateSecret(user_id int, secret models.Secret) (models.Secret, error)
	UpdateSecret(user_id int, secret models.Secret) (models.Secret, error)
	DeleteSecret(user_id int, secret_id int) error
}

type Service struct {
	Authorization
	Secret
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Secret:        NewSecretService(repos.Secret),
	}
}
