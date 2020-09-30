package repositories

import (
	"time"
	"todo/app/models"

	"github.com/jinzhu/gorm"
)

// AuthRepository godoc
type AuthRepository struct {
	Db *gorm.DB
}

// InitAuth godoc
func InitAuth(db *gorm.DB) *AuthRepository {
	return &AuthRepository{Db: db}
}

// GetSession godoc
func (repo *AuthRepository) GetSession(token string, auth *models.Auth) error {
	now := time.Now()
	return repo.Db.Where("token = ? and expire_time > ?", token, now).First(auth).Error
}
