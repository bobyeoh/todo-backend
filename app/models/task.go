package models

import "github.com/jinzhu/gorm"

// Task godoc
type Task struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(100)"`
	UserID   uint
	ColumnID uint
	User     User   `gorm:"foreignkey:UserID"`
	Column   Column `gorm:"foreignkey:ColumnID"`
}
