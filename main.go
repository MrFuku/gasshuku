package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello hot reload!!!")
	}
}

func main() {
	e := echo.New()

	e.GET("/hello", hello())

	e.Start(":8080")
}
