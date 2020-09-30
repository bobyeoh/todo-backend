package controllers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
	"todo/app"
	"todo/app/models"
	"todo/app/utils"

	"github.com/stretchr/testify/assert"
)

func TestGetColumns(t *testing.T) {
	server := app.NewServer()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := server.Echo.NewContext(req, rec)
	h := &ColumnHandler{server}
	now := time.Now()
	duration, _ := time.ParseDuration(os.Getenv("SESSION_EFFECTIVE_DURATION"))
	expireTime := now.Add(duration)
	token := utils.GetGUID()
	auth := models.Auth{
		UserID:     1,
		Token:      token,
		ExpireTime: expireTime,
	}
	c.Set("auth", &auth)

	// Assertions
	if assert.NoError(t, h.GetColumns(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
