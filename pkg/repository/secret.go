package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/qudecim/password-manager-backend/pkg/models"
)

type SecretRepository struct {
	db *sqlx.DB
}

func NewSecretRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) GetSecrets(user_id int) ([]models.Secret, error) {

	var secrets []models.Secret

	rows, err := r.db.Query("SELECT id, enycription FROM secrets WHERE user_id", user_id)
	for rows.Next() {
		var secret models.Secret

		err = rows.Scan(&secret.Id, &secret.Enycryption)
	}

	err = rows.Err()

	return secrets, err

}
