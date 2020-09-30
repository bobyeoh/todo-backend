package requests

// LoginRequest godoc
type LoginRequest struct {
	Name     string `json:"name" validate:"required" example:"bob"`
	Password string `json:"password" validate:"required" example:"11111111"`
}
