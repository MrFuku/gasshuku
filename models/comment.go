package models

import (
	"net/http"

	"github.com/MrFuku/gasshuku/auth"
	"github.com/labstack/echo"
)

type Comment struct {
	Id      int    `json:id`
	UserId  int    `json:userId`
	Content string `json:content`
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

func GetCommentAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		comments := []Comment{}
		db.Find(&comments)
		return c.JSON(http.StatusOK, comments)
	}
}
