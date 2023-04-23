package main

import (
	"echo-template/pkg/controller"
	"echo-template/pkg/repository"
	"echo-template/pkg/service"

	"github.com/labstack/echo/v4"
)

func main() {
	db := repository.NewPsqlDB()

	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	ctrl := controller.NewCtrl(svc)

	e := echo.New()
	// e.GET("/users", ctrl.GetUsers)
	e.GET("/user/:id", ctrl.GetUser)
	e.POST("/user", ctrl.CreateUser)

	e.Logger.Fatal(e.Start(":8080"))
}
