package server

import (
	"io"
	"text/template"

	"github.com/imthaghost/configtree/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// Routes will setup all the routes and templating
func (s *Server) Routes() {
	// create a new handler and dependency inject the db
	r := handlers.New(s.db)
	// Template renderer
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	// Register middleware
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.DELETE, echo.PATCH},
	}))
	// Default template renderer
	s.e.Renderer = t
	// static content
	s.e.Static("/", "assets")
	// Home page
	s.e.GET("/", r.Home)
	// login
	// s.e.GET("/login")
	// signup
	//s.e.GET("/signup")
}
