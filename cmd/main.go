package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Logger.Fatal(e.Start(":8080"))

	e.GET("/cats/:data", GetCats)
	e.POST("/cats", AddCat)
}

func GetCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	return c.JSON(http.StatusOK, map[string]string{
		"name": catName,
		"type": catType})
}

func AddCat(c echo.Context) error {
	type Cat struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}
	cat := Cat{}
	defer c.Request().Body.Close()

	if err := json.NewDecoder(c.Request().Body).Decode(&cat); err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	log.Printf("this is yout cat %#v", cat)

	return c.JSON(http.StatusOK, "We got your Cat!!!")
}
