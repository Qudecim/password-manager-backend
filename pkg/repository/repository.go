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

type Repository struct {
	Authorization
	Secret
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		Secret:        NewSecretRepository(db),
	}
}
