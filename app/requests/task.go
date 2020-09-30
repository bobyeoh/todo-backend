package requests

// TaskRequest godoc
type TaskRequest struct {
	Name     string `json:"name" validate:"required,lte=80" example:"task1"`
	ColumnID uint   `json:"column_id" validate:"required" example:"1"`
}
