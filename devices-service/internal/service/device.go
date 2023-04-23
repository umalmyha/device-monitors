package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/umalmyha/device-monitors/devices-service/internal/model"
)

type CreateDevice struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description"`
	Latitude    float32 `json:"latitude" binding:"required,lte=90,gte=-90"`
	Longitude   float32 `json:"longitude" binding:"required,lte=90,gte=-90"`
}

type UpdateDevice struct {
	Latitude    float32 `json:"latitude" binding:"required,lte=90,gte=-90"`
	Longitude   float32 `json:"longitude" binding:"required,lte=90,gte=-90"`
	Description *string `json:"description"`
}

type DeviceRepository interface {
	FindAll(ctx context.Context, qr model.GetAllDevicesQuery) ([]*model.Device, error)
	FindByID(ctx context.Context, id uuid.UUID) (*model.Device, error)
	FindByName(ctx context.Context, name string) (*model.Device, error)
	Create(ctx context.Context, dvc *model.Device) (*model.Device, error)
	Update(ctx context.Context, dvc *model.Device) (*model.Device, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type DeviceService struct {
	deviceRepo DeviceRepository
}

func NewDeviceService(deviceRepo DeviceRepository) *DeviceService {
	return &DeviceService{deviceRepo: deviceRepo}
}

func (s *DeviceService) FindAll(ctx context.Context, qr model.GetAllDevicesQuery) ([]*model.Device, error) {
	return s.deviceRepo.FindAll(ctx, qr)
}

func (s *DeviceService) FindByID(ctx context.Context, id uuid.UUID) (*model.Device, error) {
	return s.deviceRepo.FindByID(ctx, id)
}

func (s *DeviceService) Create(ctx context.Context, nd CreateDevice) (*model.Device, error) {
	found, err := s.deviceRepo.FindByName(ctx, nd.Name)
	if err != nil {
		return nil, err
	}

	if found != nil {
		return nil, fmt.Errorf("device with name %s is already present", nd.Name)
	}

	newDvc, err := s.deviceRepo.Create(ctx, &model.Device{
		ID:          uuid.New(),
		Name:        nd.Name,
		Description: nd.Description,
		Coords: model.Coords{
			Latitude:  nd.Latitude,
			Longitude: nd.Longitude,
		},
	})
	if err != nil {
		return nil, err
	}

	return newDvc, nil
}

func (s *DeviceService) Update(ctx context.Context, id uuid.UUID, upd UpdateDevice) (*model.Device, error) {
	dvc, err := s.deviceRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if dvc == nil {
		return nil, fmt.Errorf("device with id %s doesn't exist", id)
	}

	dvc.Description = upd.Description
	dvc.Latitude = upd.Latitude
	dvc.Longitude = upd.Longitude

	updDvc, err := s.deviceRepo.Update(ctx, dvc)
	if err != nil {
		return nil, err
	}

	return updDvc, nil
}

func (s *DeviceService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.deviceRepo.Delete(ctx, id)
}
