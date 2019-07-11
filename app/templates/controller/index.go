package controller

import (
	"net/http"

	"<%=repoUrl%>/model"
	"github.com/labstack/echo"
)

// Index - main page
func (ctrl *Controller) Index(c echo.Context) error {
	params := map[string]interface{}{}
	return c.Render(http.StatusOK, "index.html", params)
}

// Login - login get
func (ctrl *Controller) Login(c echo.Context) error {
	params := map[string]interface{}{
		"CSRF": c.Get("csrf").(string),
	}
	return c.Render(http.StatusOK, "login.html", params)
}

// PostLogin - login post
func (ctrl *Controller) PostLogin(c echo.Context) error {
	user := new(model.User)
	err := c.Bind(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
