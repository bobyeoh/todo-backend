package controllers

import (
	"net/http"
	"todo/app"
	"todo/app/models"
	"todo/app/repositories"
	"todo/app/responses"
	"todo/app/utils"

	"github.com/labstack/echo/v4"
)

// ColumnHandler godoc
type ColumnHandler struct {
	server *app.Server
}

// InitColumn godoc
func InitColumn(server *app.Server) *ColumnHandler {
	return &ColumnHandler{server: server}
}

// GetColumns godoc
// @Tags Column
// @Summary Get All Columns
// @Description Get All Columns
// @Success 200 {object} responses.GetColumns
// @Failure 500 {object} utils.ErrorCode "UnknownError"
// @Router /column [get]
func (handler *ColumnHandler) GetColumns(c echo.Context) error {
	var columns []responses.Column
	auth := c.Get("auth").(*models.Auth)
	// Get parameters
	columnRepo := repositories.InitColumn(handler.server.DB)
	if columnRepo.GetColumns(auth.UserID, &columns) != nil {
		return c.JSON(http.StatusInternalServerError, utils.UnknownError)
	}
	result := responses.NewGetColumns(columns)
	return c.JSON(http.StatusOK, result)
}
