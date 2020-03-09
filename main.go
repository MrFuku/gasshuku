package main

import (
	"net/http"

	"github.com/MrFuku/gasshuku/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Hoge() string {
	return "Hello Hoge!!!"
}

func hoge() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.FormValue("name")
		user := models.User{
			Name: name,
		}
		user.Save()
		return c.JSON(http.StatusOK, user)
	}
}

func hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		users := []models.User{
			{Id: 1, Name: "test1"},
			{Id: 2, Name: "test2"},
		}
		return c.JSON(http.StatusOK, users)
	}
}

func main() {
	models.InitDB()
	e := echo.New()

	e.Use(middleware.CORS())
	e.GET("/hello", hello())
	e.POST("/user", models.UserCreate())
	e.GET("/comment", models.GetCommentAll())
	e.POST("/comment", models.CommentCreate())
	e.POST("/comment/:id", models.CommentUpdate())
	e.DELETE("/comment/:id", models.CommentDelete())

	e.Start(":8080")
}
