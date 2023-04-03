package service

import (
	"github.com/qudecim/password-manager-backend/internal/models"
	"github.com/qudecim/password-manager-backend/internal/models/mysql"
)

type ConfirmationStringResponse struct {
	ConfirmationString []byte
}

func Auth(user *models.User) error {

	id, err := mysql.UserAuth(user)
	user.ID = id

	return err
}

func Register(user *models.User) (bool, error) {

	success, err := mysql.UserAdd(user)

	return success, err
}

func checkToken() {

}

func CreateToken(user *models.User) (string, error) {

}
