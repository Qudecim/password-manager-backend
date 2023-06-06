package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/qudecim/password-manager-backend/pkg/models"
)

type DeviceRepository struct {
	db *sqlx.DB
}

func NewDeviceRepository(db *sqlx.DB) *DeviceRepository {
	return &DeviceRepository{db: db}
}

func (r *DeviceRepository) GetByUid(device_uid string) (models.Device, error) {
	var device models.Device

	row := r.db.QueryRow("SELECT id, uid, name, public_key FROM devices WHERE uid = ?", device_uid)
	err := row.Scan(&device.Id, &device.Uid, &device.Name, &device.PublicKey)

	return device, err
}

func (r *DeviceRepository) CreateDevice(device models.Device) (int, error) {
	result, err := r.db.NamedExec("INSERT INTO devices (uid, name, public_key) values (:uid, :name, :public_key)", device)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
