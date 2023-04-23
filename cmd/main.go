package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Logger.Fatal(e.Start(":8080"))

	e.GET("/user/:data", GetUser)
	e.POST("/user", CreateUser)
}

func GetUser(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	return c.JSON(http.StatusOK, map[string]string{
		"name": catName,
		"type": catType})
}

type User struct {
	Name     string `json:"name"`
	Email    string `json:"type"`
	Password string `json:"password"`
}

func CreateUser(c echo.Context) error {
	var user User

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	return c.JSON(http.StatusOK, user)
}
