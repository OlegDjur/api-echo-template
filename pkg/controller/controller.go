package controller

import (
	"echo-template/pkg/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ctrl struct {
	svc service.Service
}

func (c *ctrl) CreateUser(ctx echo.Context) error {
	var request domain.UserRequest

	if err := ctx.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	result, err := c.svc.CreateUser(ctx.Request().Context(), request)

	response := domain.UserResponse{
		Name:  result.Name,
		Email: result.Email,
	}

	return ctx.JSON(http.StatusOK, response)
}

func GetUser(c echo.Context) error {
	catName := c.Param("id")

	return c.JSON(http.StatusOK, map[string]string{
		"name": catName,
		"type": catType})
}
