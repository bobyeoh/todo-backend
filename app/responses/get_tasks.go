package responses

// GetTasks godoc
type GetTasks struct {
	Tasks []Task `json:"tasks"`
}

// Task godoc
type Task struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	UserID   uint   `json:"user_id"`
	ColumnID uint   `json:"column_id"`
}

// NewGetTasks godoc
func NewGetTasks(tasks []Task) *GetTasks {
	return &GetTasks{
		Tasks: tasks,
	}
}
