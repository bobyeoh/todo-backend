package responses

// Login godoc
type Login struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// NewLogin godoc
func NewLogin(id uint, name string) *Login {
	return &Login{
		ID:   id,
		Name: name,
	}
}
