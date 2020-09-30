package services

import (
	"todo/app/models"

	"github.com/jinzhu/gorm"
)

// ColumnService godoc
type ColumnService struct {
	Db *gorm.DB
}

// InitColumn godoc
func InitColumn(db *gorm.DB) *ColumnService {
	return &ColumnService{Db: db}
}

// Create godoc
func (service *ColumnService) Create(column *models.Column) error {
	return service.Db.Create(column).Error
}

// Delete godoc
func (service *ColumnService) Delete(column *models.Column) error {
	return service.Db.Delete(column).Error
}
