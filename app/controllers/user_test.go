package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todo/app"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	loginJSON := `{"name":"test","password":"testtest"}`
	resultJSON := `{"id":1,"name":"test"}`
	server := app.NewServer()
	req := httptest.NewRequest(http.MethodPost, "/api/user/login", strings.NewReader(loginJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := server.Echo.NewContext(req, rec)
	h := &UserHandler{server}
	// Assertions
	if assert.NoError(t, h.Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, resultJSON, strings.TrimSpace(rec.Body.String()))
	}
}

func TestLoginWithoutName(t *testing.T) {
	loginJSON := `{"name":"","password":"testtest"}`
	resultJSON := `{"Code":20002,"Message":"The name is required."}`
	server := app.NewServer()
	req := httptest.NewRequest(http.MethodPost, "/api/user/login", strings.NewReader(loginJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := server.Echo.NewContext(req, rec)
	h := &UserHandler{server}
	// Assertions
	if assert.NoError(t, h.Login(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, resultJSON, strings.TrimSpace(rec.Body.String()))
	}
}

func TestLoginWithoutPassword(t *testing.T) {
	loginJSON := `{"name":"test","password":""}`
	resultJSON := `{"Code":20003,"Message":"The password is required."}`
	server := app.NewServer()
	req := httptest.NewRequest(http.MethodPost, "/api/user/login", strings.NewReader(loginJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := server.Echo.NewContext(req, rec)
	h := &UserHandler{server}
	// Assertions
	if assert.NoError(t, h.Login(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, resultJSON, strings.TrimSpace(rec.Body.String()))
	}
}

func TestLoginWithInvalidName(t *testing.T) {
	loginJSON := `{"name":"test1","password":"testtest"}`
	resultJSON := `{"Code":20001,"Message":"Invalid credentials."}`
	server := app.NewServer()
	req := httptest.NewRequest(http.MethodPost, "/api/user/login", strings.NewReader(loginJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := server.Echo.NewContext(req, rec)
	h := &UserHandler{server}
	// Assertions
	if assert.NoError(t, h.Login(c)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Equal(t, resultJSON, strings.TrimSpace(rec.Body.String()))
	}
}

func TestLoginWithoutInvalidPassword(t *testing.T) {
	loginJSON := `{"name":"test","password":"testtest1"}`
	resultJSON := `{"Code":20001,"Message":"Invalid credentials."}`
	server := app.NewServer()
	req := httptest.NewRequest(http.MethodPost, "/api/user/login", strings.NewReader(loginJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := server.Echo.NewContext(req, rec)
	h := &UserHandler{server}
	// Assertions
	if assert.NoError(t, h.Login(c)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Equal(t, resultJSON, strings.TrimSpace(rec.Body.String()))
	}
}
