package controller

import (
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	e := echo.New()
	e.Renderer = RendererTest()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	ctrl := NewController(e)
	err := ctrl.Index(ctx)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestLogin(t *testing.T) {
	e := echo.New()
	e.Renderer = RendererTest()

	req := httptest.NewRequest(http.MethodGet, "/login", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.Set("csrf", "test")

	ctrl := NewController(e)
	err := ctrl.Login(ctx)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestPostLogin(t *testing.T) {
	e := echo.New()
	payload := `{"name":"john","password":"1234567890"}`

	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	ctrl := NewController(e)
	err := ctrl.PostLogin(ctx)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, payload, strings.Trim(rec.Body.String(), "\n"))
	}
}
