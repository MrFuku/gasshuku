package models

import (
	"fmt"
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
		fmt.Println(comment, id, content)
		comment.ContentUpdate(content)
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
