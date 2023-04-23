package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/umalmyha/device-monitors/devices-service/internal/model"
	"github.com/umalmyha/device-monitors/devices-service/internal/query"
)

type MySqlDeviceRepository struct {
	db *gorm.DB
}

func NewMySqlDeviceRepository(db *gorm.DB) *MySqlDeviceRepository {
	return &MySqlDeviceRepository{db: db}
}

func (r *MySqlDeviceRepository) FindByName(ctx context.Context, name string) (*model.Device, error) {
	dvc, err := query.Device.WithContext(ctx).FindByName(name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("MySqlDeviceRepository - FindByName: %w", err)
	}
	return dvc, nil
}

func (r *MySqlDeviceRepository) FindAll(ctx context.Context, qr model.GetAllDevicesQuery) ([]*model.Device, error) {
	devices, err := query.Device.WithContext(ctx).FindAll(qr)
	if err != nil {
		return nil, fmt.Errorf("MySqlDeviceRepository - FindAll: %w", err)
	}
	return devices, nil
}

func (r *MySqlDeviceRepository) FindByID(ctx context.Context, id uuid.UUID) (*model.Device, error) {
	dvc, err := query.Device.WithContext(ctx).Where(query.Device.ID.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("MySqlDeviceRepository - FindByID: %w", err)
	}
	return dvc, nil
}

func (r *MySqlDeviceRepository) Create(ctx context.Context, dvc *model.Device) (*model.Device, error) {
	if err := query.Device.WithContext(ctx).Create(dvc); err != nil {
		return nil, fmt.Errorf("MySqlDeviceRepository - Create: %w", err)
	}
	return dvc, nil
}

func (r *MySqlDeviceRepository) Update(ctx context.Context, dvc *model.Device) (*model.Device, error) {
	if _, err := query.Device.WithContext(ctx).Updates(dvc); err != nil {
		return nil, fmt.Errorf("MySqlDeviceRepository - Update: %w", err)
	}
	return dvc, nil
}

func (r *MySqlDeviceRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if _, err := query.Device.WithContext(ctx).Where(query.Device.ID.Eq(id)).Delete(); err != nil {
		return fmt.Errorf("MySqlDeviceRepository - Delete: %w", err)
	}
	return nil
}
