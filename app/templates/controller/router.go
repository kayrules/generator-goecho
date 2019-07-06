package controller

import (
	"github.com/labstack/echo"
)

// Controller struct
type Controller struct {
}

// NewController creates and returns a new User
func NewController(e *echo.Echo) *Controller {
	ctrl := &Controller{}

	// verify demo
	e.GET("/", ctrl.Index)

	return ctrl
}
