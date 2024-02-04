package listener

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Init(e *echo.Echo) {
	setTemplates(e)
	registerHandlers(e)
}

func registerHandlers(e *echo.Echo) {
	e.GET("/", index)
}

func setTemplates(e *echo.Echo) {
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = t
}

func index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", "Hello Rana8")
}
