package database

import (
	"todo/app/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// InitDB godoc
func InitDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		panic(err.Error())
	}
	db.LogMode(true)
	Build(db)
	return db
}

// Build godoc
func Build(DB *gorm.DB) {
	user := models.User{
		Name:     "test",
		Password: "d85832fb8ce35109779e615643698d30",
	}
	if !DB.HasTable(&models.User{}) {
		DB.AutoMigrate(&models.User{})
		DB.Create(&user)
	}
	if !DB.HasTable(&models.Auth{}) {
		DB.AutoMigrate(&models.Auth{})
	}
	if !DB.HasTable(&models.Column{}) {
		DB.AutoMigrate(&models.Column{})
		if user.ID != 0 {
			backlogs := models.Column{
				Name:   "Backlog",
				UserID: user.ID,
			}
			todo := models.Column{
				Name:   "To Do",
				UserID: user.ID,
			}
			ongoing := models.Column{
				Name:   "Ongoing",
				UserID: user.ID,
			}
			done := models.Column{
				Name:   "Done",
				UserID: user.ID,
			}
			DB.Create(&backlogs)
			DB.Create(&todo)
			DB.Create(&ongoing)
			DB.Create(&done)
		}
	}
	if !DB.HasTable(&models.Task{}) {
		DB.AutoMigrate(&models.Task{})
		task := models.Task{
			Name:     "test task",
			UserID:   1,
			ColumnID: 1,
		}
		DB.Create(&task)
	}
}
