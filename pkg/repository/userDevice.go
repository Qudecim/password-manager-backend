package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/qudecim/password-manager-backend/pkg/models"
)

type UserDeviceRepository struct {
	db *sqlx.DB
}

func NewUserDeviceRepository(db *sqlx.DB) *UserDeviceRepository {
	return &UserDeviceRepository{db: db}
}

func (r *UserDeviceRepository) Get(user_id int) ([]models.Device, error) {
	var devices []models.Device

	rows, err := r.db.Query("SELECT d.id, d.key, d.name, d.public_key FROM user_device as ud LEFT JOIN device as d ON ud.device_id = d.id WHERE ud.user_id = ?", user_id)
	for rows.Next() {
		var device models.Device

		err = rows.Scan(&device.Id, &device.Uid, &device.Name, &device.PublicKey)

		devices = append(devices, device)
	}
	err = rows.Err()

	return devices, err
}

func (r *UserDeviceRepository) Add(user_id int, device models.Device) error {
	_, err := r.db.Exec("INSERT INTO user_device (user_id, device_id) values (?, ?)", user_id, device.Id)
	return err

}

func (r *UserDeviceRepository) Delete(user_id int, device_id int) error {
	_, err := r.db.Exec("DELETE FROM user_device WHERE user_id = ? AND device_id = ?", user_id, device_id)
	return err
}
