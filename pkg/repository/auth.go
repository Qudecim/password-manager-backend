package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/qudecim/password-manager-backend/pkg/models"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user models.User) (int64, error) {

	result, err := r.db.NamedExec("INSERT INTO users (name, email, password) values (:Name, :Email, :password)", user)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
