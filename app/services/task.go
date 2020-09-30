package services

import (
	"todo/app/models"
	"todo/app/requests"

	"github.com/jinzhu/gorm"
)

// TaskService godoc
type TaskService struct {
	Db *gorm.DB
}

// InitTask godoc
func InitTask(db *gorm.DB) *TaskService {
	return &TaskService{Db: db}
}

// Create godoc
func (service *TaskService) Create(task *models.Task) error {
	return service.Db.Create(task).Error
}

// Delete godoc
func (service *TaskService) Delete(task *models.Task) error {
	return service.Db.Delete(task).Error
}

// Update godoc
func (service *TaskService) Update(task *models.Task, request *requests.TaskRequest) error {
	task.ColumnID = request.ColumnID
	task.Name = request.Name
	return service.Db.Save(task).Error
}
