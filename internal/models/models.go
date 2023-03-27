package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: подходящей записи не найдено")

type User struct {
	ID                 int
	Email              string
	PublicKey          string
	ConfirmationString string
}
