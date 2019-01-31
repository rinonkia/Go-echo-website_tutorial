package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var templates map[string]*template.Template

type Template struct {
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return templates[name].ExecuteTemplate(w, "layout.html", data)
}

func main() {
	e := echo.New()

	t := &Template{}
	e.Renderer = t

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/public/css/", "./public/css/")
	e.Static("/public/js/", "./poblic/js/")
	e.Static("/public/img/", "./public/img/")

	e.GET("/", HandleIndexGet)
	e.GET("/api/hello", HandleAPIHelloGet)

	e.Logger.Fatal(e.Start(":3000"))
}

func init() {
	loadTemplates()
}

func loadTemplates() {
	var baseTemplate = "template/layout.html"
	templates = make(map[string]*template.Template)
	templates["index"] = template.Must(
		template.ParseFiles(baseTemplate, "template/hello.html"))
}
func HandleIndexGet(c echo.Context) error {
	return c.Render(http.StatusOK, "index", "World")
}

func HandleAPIHelloGet(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"Hello": "world"})
}
