package service

import (
	"github.com/google/uuid"
	"github.com/qudecim/password-manager-backend/pkg/models"
	"github.com/qudecim/password-manager-backend/pkg/repository"
)

type DeviceService struct {
	repo       repository.UserDevice
	repoDevice repository.Device
}

func NewDeviceService(repo repository.UserDevice, repoDevice repository.Device) *DeviceService {
	return &DeviceService{repo: repo, repoDevice: repoDevice}
}

func (s *DeviceService) CreateDevice(device models.Device) (models.Device, error) {

	device.Uid = uuid.New().String()

	id, err := s.repoDevice.CreateDevice(device)
	if err != nil {
		return device, err
	}

	device.Id = id

	return device, nil
}

func (s *DeviceService) ConnectDevice(uid string) (models.Device, error) {

	device, err := s.repoDevice.GetByUid(uid)
	if err != nil {
		return device, err
	}

	return device, nil
}
