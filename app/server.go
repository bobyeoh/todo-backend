package app

import (
	"os"
	"todo/app/database"
	"todo/app/validation"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server godoc
type Server struct {
	Echo *echo.Echo
	DB   *gorm.DB
}

// NewServer godoc
func NewServer() *Server {
	e := echo.New()
	e.Validator = &validation.CustomValidator{Validator: validator.New()}
	if os.Getenv("ENV") == "dev" {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowCredentials: true,
			AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
			AllowOrigins:     []string{"*"},
		}))
	}
	return &Server{
		Echo: e,
		DB:   database.InitDB(),
	}
}

// Start godoc
func (server *Server) Start(addr string) error {
	return server.Echo.Start(":" + addr)
}
