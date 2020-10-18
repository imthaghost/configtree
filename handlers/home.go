package handlers

import (
	"log"
	"net/http"

	"github.com/imthaghost/configtree/models"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Home(c echo.Context) error {
	// log visitors
	visitor := models.CreateVisitor(c.RealIP(), c.Request().RemoteAddr)
	log.Println(visitor)
	// render index page
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
}
