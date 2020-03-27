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
	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = t
	// // static files
	// e.Static("/", "assets")
	// // template render
	// renderer := &TemplateRenderer{
	// 	templates: template.Must(template.ParseGlob("*.html")),
	// }
	// e.Renderer = renderer
	// // Named route "index"
	// e.GET("/", func(c echo.Context) error {
	// 	tracks := api.Track()
	// 	fmt.Println(tracks)
	// 	return c.Render(http.StatusOK, "index.html", map[string]interface{}{

	// 		"name":       "Ransom - Lil Tecca & Juice Wrld",
	// 		"audioURL":   "http://localhost:5000/music/ransom.mp3",
	// 		"artworkURL": "https://content-images.p-cdn.com/images/public/int/7/8/3/3/00602508243387_1080W_1080H.jpg",
	// 	})

	// }).Name = "index"
	e.GET("/", controllers.Root)

	e.Logger.Fatal(e.Start(":5000"))

}
