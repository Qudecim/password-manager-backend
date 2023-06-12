package service

import (
	"github.com/qudecim/password-manager-backend/pkg/models"
	"github.com/qudecim/password-manager-backend/pkg/repository"
)

type UserDeviceService struct {
	repo       repository.UserDevice
	repoDevice repository.Device
}

func NewUserDeviceService(repo repository.UserDevice, repoDevice repository.Device) *UserDeviceService {
	return &UserDeviceService{repo: repo, repoDevice: repoDevice}
}

func (s *UserDeviceService) GetUserDevice(user_id int) ([]models.Device, error) {
	return s.repo.Get(user_id)
}

func (s *UserDeviceService) AddUserDevice(user_id int, device_uid string) (*models.Device, error) {

	device, err := s.repoDevice.GetByUid(device_uid)
	if err != nil {
		return nil, err
	}

	s.repo.Add(user_id, device)
	if err != nil {
		return nil, err
	}

	return &device, nil
}

func (s *UserDeviceService) DeleteUserDevice(user_id int, device_id int) error {
	return s.repo.Delete(user_id, device_id)
}
