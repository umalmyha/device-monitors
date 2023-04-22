package repository

import "gorm.io/gorm"

type MySqlDeviceRepository struct {
	db *gorm.DB
}
