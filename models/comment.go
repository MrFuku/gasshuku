package models

import (
	"net/http"
	"strconv"

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

func (comment *Comment) ContentUpdate(value string) {
	db.Model(&comment).Update("content", value)
}

func (comment *Comment) CommentDelete() {
	db.Delete(&comment)
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

func CommentUpdate() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		content := c.FormValue("content")
		comment := Comment{Id: id}
		comment.ContentUpdate(content)
		return c.JSON(http.StatusOK, comment)
	}
}

func CommentDelete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		comment := Comment{Id: id}
		comment.CommentDelete()
		return nil
	}
}

func GetCommentAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		comments := []Comment{}
		db.Find(&comments)
		return c.JSON(http.StatusOK, comments)
	}
}
