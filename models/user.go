package models

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Id   int    `json:id`
	Name string `json:name`
}

func (user *User) Save() {
	fmt.Println(user)
	db.Create(user)
}

func UserCreate() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.FormValue("name")
		user := User{
			Name: name,
		}
		user.Save()
		return c.JSON(http.StatusOK, user)
	}
}
