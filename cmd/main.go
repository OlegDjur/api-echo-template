package main

import (
	"echo-template/pkg/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Logger.Fatal(e.Start(":8080"))

	e.GET("/user/:data", controller.GetUser)
	e.POST("/user", controller.CreateUser)
}
