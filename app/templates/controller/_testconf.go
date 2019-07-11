package controller

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
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
