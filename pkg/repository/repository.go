package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Secret interface {
}

type Repository struct {
	Authorization
	Secret
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
