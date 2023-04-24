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

func (s *SecretService) GetSecrets(user_id int) ([]models.Secret, error) {
	return s.repo.GetSecrets(user_id)
}

func (s *SecretService) GetSecret(user_id int, secret_id int) (models.Secret, error) {
	return s.repo.GetSecret(user_id, secret_id)
}

func (s *SecretService) CreateSecret(user_id int, secret models.Secret) (int, error) {
	return s.repo.CreateSecret(user_id, secret)
}

func (s *SecretService) UpdateSecret(user_id int, secret models.Secret) error {
	return s.repo.UpdateSecret(user_id, secret)
}

func (s *SecretService) DeleteSecret(user_id int, secret_id int) error {
	return s.repo.DeleteSecret(user_id, secret_id)
}
