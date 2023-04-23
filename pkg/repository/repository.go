package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/qudecim/password-manager-backend/pkg/models"
)

type Authorization interface {
	CreateUser(user models.User) (int64, error)
	GetUserByIdAndPassword(user models.User) (models.User, error)
}

type Secret interface {
	GetSecrets(user_id int) ([]models.Secret, error)
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
