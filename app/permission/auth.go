package permission

import (
	"net/http"
	"os"
	"time"
	"todo/app"
	"todo/app/models"
	"todo/app/repositories"
	"todo/app/services"
	"todo/app/utils"

	"github.com/labstack/echo/v4"
)

// Auth godoc
func Auth(server *app.Server) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := models.Auth{}
			token, _ := c.Cookie(os.Getenv("COOKIE_KEY"))
			if token == nil {
				return c.JSON(http.StatusForbidden, utils.AccessDenied)
			}
			authRepo := repositories.InitAuth(server.DB)
			authService := services.InitAuth(server.DB)
			// init repo and service
			if authRepo.GetSession(token.Value, &auth) != nil {
				return c.JSON(http.StatusForbidden, utils.AccessDenied)
			}
			if auth.UserID == 0 {
				return c.JSON(http.StatusForbidden, utils.AccessDenied)
			}
			// verify token
			now := time.Now()
			duration, _ := time.ParseDuration(os.Getenv("SESSION_EFFECTIVE_DURATION"))
			expireTime := now.Add(duration)
			authService.ExtendKey(&auth, expireTime)
			// renew token
			c.Set("auth", &auth)
			return next(c)
		}
	}
}
