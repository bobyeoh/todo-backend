package responses

import "todo/app/models"

// NewCreateTask godoc
func NewCreateTask(task *models.Task) *Task {
	return &Task{
		ID:       task.ID,
		Name:     task.Name,
		UserID:   task.UserID,
		ColumnID: task.ColumnID,
	}
}
