package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// Index - login get
func (ctrl *Controller) Index(c echo.Context) error {
	params := map[string]interface{}{}
	return c.Render(http.StatusOK, "index.html", params)
}
