package main

import (
	"io"
	"text/template"

	"<%=repoUrl%>/config"
	"<%=repoUrl%>/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	}))

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("view/*.html")),
	}
	e.Renderer = renderer

	e.Static("/assets", "assets")
	controller.NewController(e)

	// Start server
	e.Logger.Fatal(e.Start(":" + config.Port))
}
