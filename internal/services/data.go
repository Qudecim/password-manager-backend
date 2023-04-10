package service

import (
	"github.com/qudecim/password-manager-backend/internal/models"
	"github.com/qudecim/password-manager-backend/internal/models/mysql"
)

func DataAdd(user *models.User, data string) error {

	err := mysql.DataAdd(user, data)

	return err
}
