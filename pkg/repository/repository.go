package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/qudecim/password-manager-backend/pkg/models"
)

type Authorization interface {
	CreateUser(user models.User) (int64, error)
}

type Secret interface {
}

type Repository struct {
	Authorization
	Secret
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}
