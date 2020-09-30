package services

import (
	"time"
	"todo/app/models"

	"github.com/jinzhu/gorm"
)

// UserService godoc
type UserService struct {
	Db *gorm.DB
}

// InitUser godoc
func InitUser(db *gorm.DB) *UserService {
	return &UserService{Db: db}
}

// Update godoc
func (service *UserService) Update(user *models.User) error {
	return service.Db.Save(&user).Error
}

// AddRetry godoc
func (service *UserService) AddRetry(user *models.User) error {
	return service.Db.Model(&user).UpdateColumn("retry", gorm.Expr("retry + 1")).Error
}

// Lock godoc
func (service *UserService) Lock(user *models.User) error {
	user.Retry = 0
	user.LockTime = time.Now()
	return service.Db.Save(&user).Error
}
