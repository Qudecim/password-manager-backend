package service

import (
	"github.com/qudecim/password-manager-backend/internal/models"
	"github.com/qudecim/password-manager-backend/internal/models/mysql"
)

func DataAdd(user *models.User, enycryptedData string) error {

	data := models.Data{UserId: user.ID, Data: enycryptedData}

	err := mysql.DataAdd(data)

	return err
}

func DataGet(user *models.User) ([]models.Data, error) {

	return mysql.DataGet(user)

}
