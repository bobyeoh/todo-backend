package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User godoc
type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(50);unique_index"`
	Password string `json:"password" gorm:"type:varchar(32)"`
	Retry    int    `json:"retry" gorm:"type:int"`
	LockTime time.Time
}
