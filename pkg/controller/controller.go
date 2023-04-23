package controller

import (
	"echo-template/pkg/domain"
	"net/http"

	"github.com/labstack/echo"
)

func CreateUser(c echo.Context) error {
	var user domain.UserRequest

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	return c.JSON(http.StatusOK, user)
}
