package controller

import (
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

// -- renderer
type TemplateRendererTest struct {
	templates *template.Template
}

func (t *TemplateRendererTest) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

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
