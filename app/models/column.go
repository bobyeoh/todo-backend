package models

import "github.com/jinzhu/gorm"

// Column godoc
type Column struct {
	gorm.Model
	Name   string `json:"name" gorm:"type:varchar(100)"`
	UserID uint
	User   User `gorm:"foreignkey:UserID"`
}
