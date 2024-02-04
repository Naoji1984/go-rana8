package routing

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Init(e *echo.Echo) {
	e.Use(middleware.Gzip())
	setSecurityHeaders(e)
	setTemplates(e)
	registerHandlers(e)
}

func setSecurityHeaders(e *echo.Echo) {
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "SAMEORIGIN",
		HSTSMaxAge:            2592000,
		HSTSExcludeSubdomains: false,
		ContentSecurityPolicy: "default-src 'none'; script-src-elem https://cdn.jsdelivr.net/; style-src-elem https://cdn.jsdelivr.net/;",
	}))
	e.Use(setAdditionalSecurityHeaders())
}

func setAdditionalSecurityHeaders() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			res := c.Response()
			res.Header().Set("Referrer-Policy", "no-referrer")
			res.Header().Set("Permissions-Policy", "geolocation=(),microphone=(),camera=()")
			return next(c)
		}
	}
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
