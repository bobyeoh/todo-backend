package responses

// GetColumns godoc
type GetColumns struct {
	Columns []Column `json:"columns"`
}

// Column godoc
type Column struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	UserID uint   `json:"user_id"`
}

// NewGetColumns godoc
func NewGetColumns(columns []Column) *GetColumns {
	return &GetColumns{
		Columns: columns,
	}
}
