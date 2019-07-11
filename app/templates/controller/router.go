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

	// index demo
	e.GET("/", ctrl.Index)
	e.GET("/login", ctrl.Login)
	e.POST("/login", ctrl.PostLogin)

	return ctrl
}
