package service

import (
	"github.com/qudecim/password-manager-backend/pkg/models"
	"github.com/qudecim/password-manager-backend/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int64, error)
}

type Secret interface {
}

type Service struct {
	Authorization
	Secret
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
