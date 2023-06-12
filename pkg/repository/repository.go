package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/qudecim/password-manager-backend/pkg/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUserByIdAndPassword(user models.User) (models.User, error)
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
	Add(user_id int, secret models.Device) error
	Delete(user_id int, device_id int) error
}

type Device interface {
	GetByUid(device_key string) (models.Device, error)
	CreateDevice(device models.Device) (int, error)
}

type Repository struct {
	Authorization
	Secret
	UserDevice
	Device
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		Secret:        NewSecretRepository(db),
		UserDevice:    NewUserDeviceRepository(db),
		Device:        NewDeviceRepository(db),
	}
}
