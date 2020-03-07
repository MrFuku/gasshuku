package models

import (
	"net/http"

	"github.com/MrFuku/gasshuku/auth"
	"github.com/labstack/echo"
)

type Comment struct {
	Id      int    `json:id`
	UserId  int    `json:id`
	Content string `json:name`
}

func (comment *Comment) Save() {
	db.Create(comment)
}

func CommentCreate() echo.HandlerFunc {
	return func(c echo.Context) error {
		content := c.FormValue("content")
		comment := Comment{
			UserId:  auth.LoginUser(c),
			Content: content,
		}
		comment.Save()
		return c.JSON(http.StatusOK, comment)
	}
}
