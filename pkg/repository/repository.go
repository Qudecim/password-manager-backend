package repository

type Authorization interface {
}

type Secret interface {
}

type Repository struct {
	Authorization
	Secret
}

func NewRepository() *Repository {
	return &Repository{}
}
