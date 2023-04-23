package controller

import (
	"echo-template/pkg/domain"
	"echo-template/pkg/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ctrl struct {
	svc service.Service
}

func NewCtrl(svc service.Service) *ctrl {
	return &ctrl{
		svc: svc,
	}
}

func (c *ctrl) CreateUser(ctx echo.Context) error {
	var request domain.UserRequest

	if err := ctx.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}

	result, err := c.svc.CreateUser(ctx.Request().Context(), request)
	if err != nil {
		fmt.Println(err)
		return err
	}

	response := domain.UserResponse{
		ID:    result.ID,
		Name:  result.Name,
		Email: result.Email,
	}

	return ctx.JSON(http.StatusOK, response)
}

func (c *ctrl) GetUser(ctx echo.Context) error {
	request := ctx.Param("id")

	requestInt, err := strconv.Atoi(request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}

	response, err := c.svc.GetUser(ctx.Request().Context(), requestInt)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return ctx.JSON(http.StatusOK, response)
}
