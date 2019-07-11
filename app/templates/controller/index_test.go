package controller

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

// TemplateRendererTest - for unit test
type TemplateRendererTest struct {
	templates *template.Template
}

// Render - for unit test
func (t *TemplateRendererTest) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// RendererTest - for unit test
func RendererTest() *TemplateRendererTest {
	renderer := &TemplateRendererTest{
		templates: template.Must(template.ParseGlob("../view/*.html")),
	}
	return renderer
}

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
