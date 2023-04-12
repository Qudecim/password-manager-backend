package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: подходящей записи не найдено")

type User struct {
	ID       int
	Email    string
	Password string
}

type Data struct {
	ID     int
	UserId int
	Data   string
}

type Device struct {
	ID        int
	UserId    int
	Name      string
	Uuid      string
	PublicKey string
}
