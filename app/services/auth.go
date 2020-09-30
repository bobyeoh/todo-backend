package services

import (
	"time"
	"todo/app/models"

	"github.com/jinzhu/gorm"
)

// AuthServices godoc
type AuthServices struct {
	Db *gorm.DB
}

// InitAuth godoc
func InitAuth(db *gorm.DB) *AuthServices {
	return &AuthServices{Db: db}
}

// SetSession godoc
func (service *AuthServices) SetSession(auth *models.Auth) error {
	return service.Db.Create(&auth).Error
}

// ExtendKey godoc
func (service *AuthServices) ExtendKey(auth *models.Auth, expireTime time.Time) error {
	auth.ExpireTime = expireTime
	return service.Db.Save(&auth).Error
}

// ClearSession godoc
func (service *AuthServices) ClearSession() error {
	now := time.Now()
	return service.Db.Where("expire_time < ?", now).Delete(&models.Auth{}).Error
}
