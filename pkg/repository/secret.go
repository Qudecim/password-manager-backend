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

	rows, err := r.db.Query("SELECT id, user_id, enycription FROM secrets WHERE user_id = ?", user_id)
	for rows.Next() {
		var secret models.Secret

		err = rows.Scan(&secret.Id, &secret.UserId, &secret.Enycryption)

		secrets = append(secrets, secret)
	}
	err = rows.Err()

	return secrets, err
}

func (r *AuthRepository) GetSecret(user_id int, secret_id int) (models.Secret, error) {
	var secret models.Secret

	row := r.db.QueryRow("SELECT id, user_id, enycription FROM secrets WHERE user_id = ? AND id = ?", user_id, secret_id)
	err := row.Scan(&secret.Id, &secret.UserId, &secret.Enycryption)

	return secret, err
}

func (r *AuthRepository) CreateSecret(user_id int, secret models.Secret) (int, error) {
	result, err := r.db.Exec("INSERT INTO secrets (user_id, enycription) values (?, ?)", user_id, secret.Enycryption)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *AuthRepository) UpdateSecret(user_id int, secret models.Secret) error {
	_, err := r.db.Exec("UPDATE secrets SET enycription = ? WHERE user_id = ?", secret.Enycryption, user_id)
	return err
}

func (r *AuthRepository) DeleteSecret(user_id int, secret_id int) error {
	_, err := r.db.Exec("DELETE FROM secrets WHERE user_id = ? AND id = ?", user_id, secret_id)
	return err
}
