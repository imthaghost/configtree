package controllers

import (
	"fmt"
	"net/http"

	"github.com/imthaghost/configtree/models"
	"github.com/labstack/echo"
)

// Root ...
func Root(c echo.Context) error {
	// visitor to page
	visitor := models.CreateVisitor(c.RealIP(), c.Request().RemoteAddr)
	// who is
	fmt.Println(visitor)
	// return html
	return c.String(http.StatusOK, "hi")

}
