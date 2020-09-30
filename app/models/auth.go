package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Auth godoc
type Auth struct {
	gorm.Model
	Token      string `gorm:"type:varchar(50);unique_index"`
	ExpireTime time.Time
	UserID     uint `gorm:"type:int"`
	User       User `gorm:"foreignkey:UserID"`
}
