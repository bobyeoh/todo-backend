package repositories

import (
	"todo/app/models"
	"todo/app/responses"

	"github.com/jinzhu/gorm"
)

// TaskRepository godoc
type TaskRepository struct {
	Db *gorm.DB
}

// InitTask godoc
func InitTask(db *gorm.DB) *TaskRepository {
	return &TaskRepository{Db: db}
}

// GetTasks godoc
func (repo *TaskRepository) GetTasks(columnID uint, userID uint, tasks *[]responses.Task) error {
	return repo.Db.Model(&models.Task{}).Where("user_id = ? and column_id = ?", userID, columnID).Scan(tasks).Error
}

// GetTask godoc
func (repo *TaskRepository) GetTask(id uint, task *models.Task) error {
	return repo.Db.Where("id = ?", id).First(task).Error
}
