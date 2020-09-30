package controllers

import (
	"net/http"
	"os"
	"time"
	"todo/app"
	"todo/app/models"
	"todo/app/repositories"
	"todo/app/requests"
	"todo/app/responses"
	"todo/app/services"
	"todo/app/utils"
	"todo/app/validation"

	"github.com/labstack/echo/v4"
)

// UserHandler godoc
type UserHandler struct {
	server *app.Server
}

// InitUser godoc
func InitUser(server *app.Server) *UserHandler {
	return &UserHandler{server: server}
}

// Login godoc
// @Tags User
// @Summary Login
// @Description Login
// @Param params body requests.LoginRequest true "Name and password"
// @Success 200 {object} responses.Login
// @Failure 400 {object} utils.ErrorCode
// @Failure 401 {object} utils.ErrorCode "InvalidCredentials"
// @Failure 500 {object} utils.ErrorCode "UnknownError"
// @Router /user/login [post]
func (handler *UserHandler) Login(c echo.Context) error {
	loginRequest := new(requests.LoginRequest)
	if err := c.Bind(loginRequest); err != nil {
		return err
	}
	// Get parameters

	if err := c.Validate(loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, validation.ProcessError(err))
	}
	// valitation

	now := time.Now()
	userRepo := repositories.InitUser(handler.server.DB)
	userService := services.InitUser(handler.server.DB)
	authService := services.InitAuth(handler.server.DB)
	// init

	user := models.User{}
	userRepo.GetUserByName(&user, loginRequest.Name)
	if user.ID == 0 || user.Password != utils.MD5(loginRequest.Password) {
		if user.ID != 0 {
			userService.AddRetry(&user)
			if user.Retry == 3 {
				userService.Lock(&user)
			}
		}
		// maximum 3 times retries, lock if exceeded
		return c.JSON(http.StatusUnauthorized, utils.InvalidCredentials)
	}
	// Verify name and password

	if now.Sub(user.LockTime).Minutes() < 10 {
		return c.JSON(http.StatusUnauthorized, utils.TooManyRetry)
	}
	// account has been locked

	duration, _ := time.ParseDuration(os.Getenv("SESSION_EFFECTIVE_DURATION"))
	expireTime := now.Add(duration)
	token := utils.GetGUID()
	auth := models.Auth{
		UserID:     user.ID,
		Token:      token,
		ExpireTime: expireTime,
	}
	if authService.SetSession(&auth) != nil {
		return c.JSON(http.StatusInternalServerError, utils.UnknownError)
	}
	authService.ClearSession() // Clear Expired Session
	// Create token

	cookie := http.Cookie{
		Name:     os.Getenv("COOKIE_KEY"),
		Value:    auth.Token,
		Expires:  expireTime,
		Path:     "/",
		HttpOnly: true,
	}
	c.SetCookie(&cookie)
	// set cookie

	user.Retry = 0
	userService.Db.Save(&user)
	// reset retry

	result := responses.NewLogin(user.ID, user.Name)
	return c.JSON(http.StatusOK, result)
}
