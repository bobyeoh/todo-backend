package repositories

import (
	"todo/app/models"

	"github.com/jinzhu/gorm"
)

// UserRepository godoc
type UserRepository struct {
	Db *gorm.DB
}

// InitUser godoc
func InitUser(db *gorm.DB) *UserRepository {
	return &UserRepository{Db: db}
}

// GetUserByName godoc
func (repo *UserRepository) GetUserByName(user *models.User, name string) {
	repo.Db.Where("name = ?", name).Find(user)
}
