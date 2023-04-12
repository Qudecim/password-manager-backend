package mysql

import (
	"context"

	"github.com/qudecim/password-manager-backend/internal/app"
	"github.com/qudecim/password-manager-backend/internal/models"
)

func DeviceList(user *models.User) ([]models.Device, error) {

	var devices []models.Device

	rows, err := app.DB.Query("select name, uuid, public_key, user_id from Devices where user_id = ?", user.ID)
	if err != nil {
		return devices, err
	}

	for rows.Next() {
		var device models.Device
		if err := rows.Scan(&device.Name, &device.Uuid, &device.PublicKey, &device.UserId); err != nil {
			return devices, err
		}
		devices = append(devices, device)
	}

	return devices, nil
}

func DeviceGet(user *models.User, deviceUuid string) (*models.Device, error) {
	var device models.Device
	err := app.DB.QueryRow("select * from Devices where uuid = ? AND user_id = ?", deviceUuid, user.ID).Scan(&device.ID, &device.Name, &device.Uuid, &device.PublicKey, &device.UserId)
	if err != nil {
		return nil, err
	}
	return &device, nil
}

func DeviceAdd(device *models.Device) error {
	_, err := app.DB.ExecContext(context.Background(), "INSERT INTO Devices (name, uuid, public_key, user_id) VALUES (?,?,?,?,?)", device.Name, device.Uuid, device.PublicKey, device.UserId)
	return err
}

func DeviceDelete(device *models.Device) error {
	_, err := app.DB.ExecContext(context.Background(), "DELETE FROM Devices WHERE uuid = ? AND user_id", device.Uuid, device.UserId)
	return err
}
