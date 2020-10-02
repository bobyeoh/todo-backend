package controllers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
	"todo/app"
	"todo/app/models"
	"todo/app/utils"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	createJSON := `{"name":"test task","column_id": 1}`
	server := app.NewServer()
	request := httptest.NewRequest(http.MethodPost, "/api/task", strings.NewReader(createJSON))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	record := httptest.NewRecorder()
	context := server.Echo.NewContext(request, record)
	handler := &TaskHandler{server}
	now := time.Now()
	duration, _ := time.ParseDuration(os.Getenv("SESSION_EFFECTIVE_DURATION"))
	expireTime := now.Add(duration)
	token := utils.GetGUID()
	auth := models.Auth{
		UserID:     1,
		Token:      token,
		ExpireTime: expireTime,
	}
	context.Set("auth", &auth)
	if assert.NoError(t, handler.CreateTask(context)) {
		assert.Equal(t, http.StatusCreated, record.Code)
	}
}

func TestCreateTaskWithoutName(t *testing.T) {
	createJSON := `{"name":"","column_id": 1}`
	resultJSON := `{"Code":30004,"Message":"The task name is required."}`
	server := app.NewServer()
	request := httptest.NewRequest(http.MethodPost, "/api/task", strings.NewReader(createJSON))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	record := httptest.NewRecorder()
	context := server.Echo.NewContext(request, record)
	handler := &TaskHandler{server}
	now := time.Now()
	duration, _ := time.ParseDuration(os.Getenv("SESSION_EFFECTIVE_DURATION"))
	expireTime := now.Add(duration)
	token := utils.GetGUID()
	auth := models.Auth{
		UserID:     1,
		Token:      token,
		ExpireTime: expireTime,
	}
	context.Set("auth", &auth)
	if assert.NoError(t, handler.CreateTask(context)) {
		assert.Equal(t, http.StatusBadRequest, record.Code)
		assert.Equal(t, resultJSON, strings.TrimSpace(record.Body.String()))
	}
}

func TestCreateTaskWithoutColumn(t *testing.T) {
	createJSON := `{"name":"asaa","column_id": 0}`
	resultJSON := `{"Code":30003,"Message":"The column id is required."}`
	server := app.NewServer()
	request := httptest.NewRequest(http.MethodPost, "/api/task", strings.NewReader(createJSON))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	record := httptest.NewRecorder()
	context := server.Echo.NewContext(request, record)
	handler := &TaskHandler{server}
	now := time.Now()
	duration, _ := time.ParseDuration(os.Getenv("SESSION_EFFECTIVE_DURATION"))
	expireTime := now.Add(duration)
	token := utils.GetGUID()
	auth := models.Auth{
		UserID:     1,
		Token:      token,
		ExpireTime: expireTime,
	}
	context.Set("auth", &auth)
	if assert.NoError(t, handler.CreateTask(context)) {
		assert.Equal(t, http.StatusBadRequest, record.Code)
		assert.Equal(t, resultJSON, strings.TrimSpace(record.Body.String()))
	}
}

func TestCreateTaskWithoutNotExistColumn(t *testing.T) {
	createJSON := `{"name":"asaa","column_id": 5}`
	resultJSON := `{"Code":30002,"Message":"The column does not exist."}`
	server := app.NewServer()
	request := httptest.NewRequest(http.MethodPost, "/api/task", strings.NewReader(createJSON))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	record := httptest.NewRecorder()
	context := server.Echo.NewContext(request, record)
	handler := &TaskHandler{server}
	now := time.Now()
	duration, _ := time.ParseDuration(os.Getenv("SESSION_EFFECTIVE_DURATION"))
	expireTime := now.Add(duration)
	token := utils.GetGUID()
	auth := models.Auth{
		UserID:     1,
		Token:      token,
		ExpireTime: expireTime,
	}
	context.Set("auth", &auth)
	if assert.NoError(t, handler.CreateTask(context)) {
		assert.Equal(t, http.StatusBadRequest, record.Code)
		assert.Equal(t, resultJSON, strings.TrimSpace(record.Body.String()))
	}
}
func TestUpdateTask(t *testing.T) {
	createJSON := `{"name":"test task1","column_id": 1}`
	server := app.NewServer()
	request := httptest.NewRequest(http.MethodPut, "/api/task/", strings.NewReader(createJSON))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	record := httptest.NewRecorder()
	context := server.Echo.NewContext(request, record)
	handler := &TaskHandler{server}
	now := time.Now()
	duration, _ := time.ParseDuration(os.Getenv("SESSION_EFFECTIVE_DURATION"))
	expireTime := now.Add(duration)
	token := utils.GetGUID()
	auth := models.Auth{
		UserID:     1,
		Token:      token,
		ExpireTime: expireTime,
	}
	context.Set("auth", &auth)
	context.SetParamNames("id")
	context.SetParamValues("1")
	if assert.NoError(t, handler.UpdateTask(context)) {
		assert.Equal(t, http.StatusOK, record.Code)
	}
}

func TestUpdateTaskNotExist(t *testing.T) {
	resultJSON := `{"Code":30001,"Message":"The task does not exist."}`
	createJSON := `{"name":"test task1","column_id": 1}`
	server := app.NewServer()
	request := httptest.NewRequest(http.MethodPut, "/api/task/", strings.NewReader(createJSON))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	record := httptest.NewRecorder()
	context := server.Echo.NewContext(request, record)
	handler := &TaskHandler{server}
	now := time.Now()
	duration, _ := time.ParseDuration(os.Getenv("SESSION_EFFECTIVE_DURATION"))
	expireTime := now.Add(duration)
	token := utils.GetGUID()
	auth := models.Auth{
		UserID:     1,
		Token:      token,
		ExpireTime: expireTime,
	}
	context.Set("auth", &auth)
	context.SetParamNames("id")
	context.SetParamValues("10")
	if assert.NoError(t, handler.UpdateTask(context)) {
		assert.Equal(t, http.StatusNotFound, record.Code)
		assert.Equal(t, resultJSON, strings.TrimSpace(record.Body.String()))
	}
}

func TestUpdateTaskWithoutName(t *testing.T) {
	resultJSON := `{"Code":30004,"Message":"The task name is required."}`
	createJSON := `{"name":"","column_id": 1}`
	server := app.NewServer()
	request := httptest.NewRequest(http.MethodPut, "/api/task/", strings.NewReader(createJSON))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	record := httptest.NewRecorder()
	context := server.Echo.NewContext(request, record)
	handler := &TaskHandler{server}
	now := time.Now()
	duration, _ := time.ParseDuration(os.Getenv("SESSION_EFFECTIVE_DURATION"))
	expireTime := now.Add(duration)
	token := utils.GetGUID()
	auth := models.Auth{
		UserID:     1,
		Token:      token,
		ExpireTime: expireTime,
	}
	context.Set("auth", &auth)
	context.SetParamNames("id")
	context.SetParamValues("1")
	if assert.NoError(t, handler.UpdateTask(context)) {
		assert.Equal(t, http.StatusBadRequest, record.Code)
		assert.Equal(t, resultJSON, strings.TrimSpace(record.Body.String()))
	}
}

func TestUpdateTaskWithoutColumn(t *testing.T) {
	resultJSON := `{"Code":30003,"Message":"The column id is required."}`
	createJSON := `{"name":"test","column_id": 0}`
	server := app.NewServer()
	request := httptest.NewRequest(http.MethodPut, "/api/task/", strings.NewReader(createJSON))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	record := httptest.NewRecorder()
	context := server.Echo.NewContext(request, record)
	handler := &TaskHandler{server}
	now := time.Now()
	duration, _ := time.ParseDuration(os.Getenv("SESSION_EFFECTIVE_DURATION"))
	expireTime := now.Add(duration)
	token := utils.GetGUID()
	auth := models.Auth{
		UserID:     1,
		Token:      token,
		ExpireTime: expireTime,
	}
	context.Set("auth", &auth)
	context.SetParamNames("id")
	context.SetParamValues("1")
	if assert.NoError(t, handler.UpdateTask(context)) {
		assert.Equal(t, http.StatusBadRequest, record.Code)
		assert.Equal(t, resultJSON, strings.TrimSpace(record.Body.String()))
	}
}

func TestDeleteTask(t *testing.T) {
	server := app.NewServer()
	request := httptest.NewRequest(http.MethodDelete, "/api/task/", nil)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	record := httptest.NewRecorder()
	context := server.Echo.NewContext(request, record)
	handler := &TaskHandler{server}
	now := time.Now()
	duration, _ := time.ParseDuration(os.Getenv("SESSION_EFFECTIVE_DURATION"))
	expireTime := now.Add(duration)
	token := utils.GetGUID()
	auth := models.Auth{
		UserID:     1,
		Token:      token,
		ExpireTime: expireTime,
	}
	context.Set("auth", &auth)
	context.SetParamNames("id")
	context.SetParamValues("1")
	if assert.NoError(t, handler.DeleteTask(context)) {
		assert.Equal(t, http.StatusOK, record.Code)
	}
}

func TestGetTasks(t *testing.T) {
	server := app.NewServer()
	request := httptest.NewRequest(http.MethodGet, "/api/task", nil)
	record := httptest.NewRecorder()
	context := server.Echo.NewContext(request, record)
	handler := &TaskHandler{server}
	now := time.Now()
	duration, _ := time.ParseDuration(os.Getenv("SESSION_EFFECTIVE_DURATION"))
	expireTime := now.Add(duration)
	token := utils.GetGUID()
	auth := models.Auth{
		UserID:     1,
		Token:      token,
		ExpireTime: expireTime,
	}
	context.Set("auth", &auth)
	context.SetParamNames("column_id")
	context.SetParamValues("1")
	// Assertions
	if assert.NoError(t, handler.GetTasks(context)) {
		assert.Equal(t, http.StatusOK, record.Code)
	}
}
