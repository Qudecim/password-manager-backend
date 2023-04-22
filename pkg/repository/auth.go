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

	result, err := r.db.NamedExec("INSERT INTO users (email, password) values (:email, :password)", user)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthRepository) GetUser(user models.User) (models.User, error) {

	row := r.db.QueryRow("SELECT id FROM users WHERE email=? AND password=?", user.Email, user.Password)
	err := row.Scan(&user.Id)

	return user, err
}
