package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// Error - login get
func Error(c echo.Context) (err error) {
	params := map[string]interface{}{}
	return c.Render(http.StatusInternalServerError, "error_400.html", params)
}
