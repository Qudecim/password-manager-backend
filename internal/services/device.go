package service

import (
	"github.com/qudecim/password-manager-backend/internal/models"
	"github.com/qudecim/password-manager-backend/internal/models/mysql"
)

func DeviceList(user *models.User) ([]models.Device, error) {

	devices, err := mysql.DeviceList(user)

	return devices, err
}

func DeviceDelete(device *models.Device) error {

	err := mysql.DeviceDelete(device)

	return err
}

func DeviceAdd(user *models.User, device *models.Device) error {

	return nil
}
