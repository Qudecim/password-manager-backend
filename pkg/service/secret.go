package service

import (
	"github.com/qudecim/password-manager-backend/pkg/models"
	"github.com/qudecim/password-manager-backend/pkg/repository"
)

type SecretService struct {
	repo repository.Secret
}

func NewSecretService(repo repository.Secret) *SecretService {
	return &SecretService{repo: repo}
}

func (s *SecretService) GetAllSecrets(user_id int) ([]models.Secret, error) {
	return s.repo.GetSecrets(user_id)
}

func (s *SecretService) GetOneSecret(user_id int, secret_id int) (models.Secret, error) {

}

func (s *SecretService) CreateSecret(user_id int, secret models.Secret) (models.Secret, error) {

}

func (s *SecretService) UpdateSecret(user_id int, secret models.Secret) (models.Secret, error) {

}

func (s *SecretService) DeleteSecret(user_id int, secret_id int) error {

	return nil
}
