package service

import "github.com/qudecim/password-manager-backend/pkg/repository"

type Authorization interface {
}

type Secret interface {
}

type Service struct {
	Authorization
	Secret
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
