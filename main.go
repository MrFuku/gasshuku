package main

import (
	"net/http"

	"github.com/MrFuku/gasshuku/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		users := []User{
			{Id: 1, Name: "test1"},
			{Id: 2, Name: "test2"},
		}
		return c.JSON(http.StatusOK, users)
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.GET("/hello", hello())

	e.Start(":8080")
}
