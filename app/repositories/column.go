package repositories

import (
	"todo/app/models"
	"todo/app/responses"

	"github.com/jinzhu/gorm"
)

// ColumnRepository godoc
type ColumnRepository struct {
	Db *gorm.DB
}

// InitColumn godoc
func InitColumn(db *gorm.DB) *ColumnRepository {
	return &ColumnRepository{Db: db}
}

// GetColumns godoc
func (repo *ColumnRepository) GetColumns(userID uint, columns *[]responses.Column) error {
	return repo.Db.Model(&models.Column{}).Where("user_id = ?", userID).Scan(columns).Error
}

// GetColumn godoc
func (repo *ColumnRepository) GetColumn(id uint, userID uint, column *models.Column) error {
	return repo.Db.Where("id = ? and user_id = ?", id, userID).First(column).Error
}
