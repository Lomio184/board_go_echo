package render

import (
	"github.com/labstack/echo/v4"
	"html/template"
    "io"
)

type TemplateRenderer struct {
    templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func CreateRenderer() *TemplateRenderer {
	return &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
}