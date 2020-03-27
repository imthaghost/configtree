package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/imthaghost/configtree/models"
	"github.com/labstack/echo"
)

// Root ...
func Root(c echo.Context) error {
	// var bodyBytes []byte
	// // request body
	// if c.Request().Body != nil {
	// 	// Read the Body content
	// 	bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	// }
	// // Restore the io.ReadCloser to its original state
	// c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	// visitor to page
	visitor := models.CreateVisitor(c.RealIP(), c.Request().RemoteAddr)
	// who is
	fmt.Println(visitor)
	// return html
	return c.Render(http.StatusOK, "index", template.HTML("<p>HTML Test</p>"))

}
