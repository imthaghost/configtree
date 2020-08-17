package main

import (
	"io"
	"text/template"

	"github.com/imthaghost/configtree/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	// Log Output
	e.Use(middleware.Logger())
	// Stream Recovery
	e.Use(middleware.Recover())

	e.GET("/", controllers.Root)

	e.Logger.Fatal(e.Start(":5000"))

}
