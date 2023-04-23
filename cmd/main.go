package main

import (
	"echo-template/pkg/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	ctrl := controller.NewCtrl(svc)

	e := echo.New()
	e.GET("/users", ctrl.GetUsers)
	e.GET("/user/:id", ctrl.GetUser)
	e.POST("/user", ctrl.CreateUser)

	e.Logger.Fatal(e.Start(":8080"))
}
