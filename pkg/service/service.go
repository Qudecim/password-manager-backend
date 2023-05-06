package service

import (
	"github.com/qudecim/password-manager-backend/pkg/models"
	"github.com/qudecim/password-manager-backend/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(user models.User) (string, error)
	ParseToken(token string) (int, error)
}

type Secret interface {
	GetSecrets(user_id int) ([]models.Secret, error)
	GetSecret(user_id int, secret_id int) (models.Secret, error)
	CreateSecret(user_id int, secret models.Secret) (int, error)
	UpdateSecret(user_id int, secret models.Secret) error
	DeleteSecret(user_id int, secret_id int) error
}

type UserDevice interface {
	Get(user_id int) ([]models.Device, error)
	Add(user_id int, device string) (*models.Device, error)
	Delete(user_id int, device_id int) error
}

type Service struct {
	Authorization
	Secret
	UserDevice
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Secret:        NewSecretService(repos.Secret),
		UserDevice:    NewUserDeviceService(repos.UserDevice, repos.Device),
	}
}
