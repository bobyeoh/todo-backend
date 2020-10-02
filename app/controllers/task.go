package controllers

import (
	"net/http"
	"strconv"
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

// TaskHandler godoc
type TaskHandler struct {
	server *app.Server
}

// InitTask godoc
func InitTask(server *app.Server) *TaskHandler {
	return &TaskHandler{server: server}
}

// GetTasks godoc
// @Tags Task
// @Summary Get tasks by column id
// @Description Get tasks by column id
// @Param id path int true "Column ID"
// @Success 200 {object} responses.GetTasks
// @Failure 500 {object} utils.ErrorCode "UnknownError"
// @Router /api/task/{column_id} [get]
func (handler *TaskHandler) GetTasks(c echo.Context) error {
	var tasks []responses.Task
	columnID, _ := strconv.Atoi(c.Param("id"))
	auth := c.Get("auth").(*models.Auth)
	// Get parameters
	taskRepo := repositories.InitTask(handler.server.DB)
	if taskRepo.GetTasks(uint(columnID), auth.UserID, &tasks) != nil {
		return c.JSON(http.StatusInternalServerError, utils.UnknownError)
	}
	result := responses.NewGetTasks(tasks)
	return c.JSON(http.StatusOK, result)
}

// CreateTask godoc
// @Tags Task
// @Summary Create task
// @Description Create task
// @Param params body requests.TaskRequest true "name and column_id"
// @Success 201 {object} responses.Task
// @Failure 400 {object} utils.ErrorCode "TheColumnDoesNotExist"
// @Failure 403 {object} utils.ErrorCode "PermissionDenied"
// @Failure 404 {object} utils.ErrorCode "TheTaskDoesNotExist"
// @Failure 500 {object} utils.ErrorCode "UnknownError"
// @Router /api/task [post]
func (handler *TaskHandler) CreateTask(c echo.Context) error {
	column := models.Column{}
	auth := c.Get("auth").(*models.Auth)
	taskRequest := new(requests.TaskRequest)
	if err := c.Bind(taskRequest); err != nil {
		return err
	}
	task := models.Task{
		Name:     taskRequest.Name,
		UserID:   auth.UserID,
		ColumnID: taskRequest.ColumnID,
	}
	// Get parameters
	if err := c.Validate(taskRequest); err != nil {
		return c.JSON(http.StatusBadRequest, validation.ProcessError(err))
	}
	// validation

	taskService := services.InitTask(handler.server.DB)
	columnRepo := repositories.InitColumn(handler.server.DB)
	// init

	if columnRepo.GetColumn(task.ColumnID, auth.UserID, &column) != nil {
		return c.JSON(http.StatusBadRequest, utils.TheColumnDoesNotExist)
	}
	if column.ID == 0 {
		return c.JSON(http.StatusBadRequest, utils.TheColumnDoesNotExist)
	}
	// verify column

	if taskService.Create(&task) != nil {
		return c.JSON(http.StatusInternalServerError, utils.UnknownError)
	}
	// create

	title := "Task Created"
	content := "Task " + task.Name + " has been created."
	go utils.SendMail("manager@email.com", title, content)
	// send email

	taskResponse := responses.NewCreateTask(&task)
	return c.JSON(http.StatusCreated, taskResponse)
}

// UpdateTask godoc
// @Tags Task
// @Summary Update Task
// @Description Update Task
// @Param id path int true "Task ID"
// @Param params body requests.TaskRequest true "name and column_id"
// @Success 200
// @Failure 400 {object} utils.ErrorCode "TheColumnDoesNotExist"
// @Failure 403 {object} utils.ErrorCode "PermissionDenied"
// @Failure 404 {object} utils.ErrorCode "TheTaskDoesNotExist"
// @Failure 500 {object} utils.ErrorCode "UnknownError"
// @Router /api/task/{id} [put]
func (handler *TaskHandler) UpdateTask(c echo.Context) error {
	var task models.Task
	var newColumn models.Column
	id, _ := strconv.Atoi(c.Param("id"))
	auth := c.Get("auth").(*models.Auth)
	taskRequest := new(requests.TaskRequest)
	if err := c.Bind(taskRequest); err != nil {
		return err
	}
	// Get parameters

	if err := c.Validate(taskRequest); err != nil {
		return c.JSON(http.StatusBadRequest, validation.ProcessError(err))
	}
	// validation
	taskService := services.InitTask(handler.server.DB)
	taskRepo := repositories.InitTask(handler.server.DB)
	columnRepo := repositories.InitColumn(handler.server.DB)
	//init

	if taskRepo.GetTask(uint(id), &task) != nil {
		return c.JSON(http.StatusNotFound, utils.TheTaskDoesNotExist)
	}
	if task.ID == 0 {
		return c.JSON(http.StatusNotFound, utils.TheTaskDoesNotExist)
	}
	// verify task

	if task.UserID != auth.UserID {
		return c.JSON(http.StatusForbidden, utils.PermissionDenied)
	}
	// verify permission

	title := "Task Update"
	content := ""
	if task.Name != taskRequest.Name {
		content = "Task " + task.Name + " updated to " + taskRequest.Name
	}
	// organize email content

	if task.ColumnID != taskRequest.ColumnID {
		if columnRepo.GetColumn(taskRequest.ColumnID, auth.UserID, &newColumn) != nil {
			return c.JSON(http.StatusBadRequest, utils.TheColumnDoesNotExist)
		}
		if newColumn.ID == 0 {
			return c.JSON(http.StatusBadRequest, utils.TheColumnDoesNotExist)
		}
		content = "Task " + task.Name + " was moved from " + task.Column.Name + "to " + newColumn.Name
	}
	// verify column

	if taskService.Update(&task, taskRequest) != nil {
		return c.JSON(http.StatusInternalServerError, utils.UnknownError)
	}
	// update task

	if content != "" {
		go utils.SendMail("manager@email.com", title, content)
	}
	// send email

	return c.NoContent(http.StatusOK)
}

// DeleteTask godoc
// @Tags Task
// @Summary Delete Task
// @Description Delete Task
// @Param id path int true "Task ID"
// @Success 200
// @Failure 403 {object} utils.ErrorCode "PermissionDenied"
// @Failure 404 {object} utils.ErrorCode "TheTaskDoesNotExist"
// @Failure 500 {object} utils.ErrorCode "UnknownError"
// @Router /api/task/{id} [delete]
func (handler *TaskHandler) DeleteTask(c echo.Context) error {
	var task models.Task
	id, _ := strconv.Atoi(c.Param("id"))
	auth := c.Get("auth").(*models.Auth)
	// Get parameters
	taskRepo := repositories.InitTask(handler.server.DB)
	taskService := services.InitTask(handler.server.DB)
	if taskRepo.GetTask(uint(id), &task) != nil {
		return c.JSON(http.StatusInternalServerError, utils.UnknownError)
	}
	if task.ID == 0 {
		return c.JSON(http.StatusNotFound, utils.TheTaskDoesNotExist)
	}
	if task.UserID != auth.UserID {
		return c.JSON(http.StatusForbidden, utils.PermissionDenied)
	}
	title := "Task Delete"
	content := "Task " + task.Name + " has been deleted."
	if taskService.Delete(&task) != nil {
		return c.JSON(http.StatusInternalServerError, utils.UnknownError)
	}
	go utils.SendMail("manager@email.com", title, content)
	return c.NoContent(http.StatusOK)
}
