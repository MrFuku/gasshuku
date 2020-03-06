package models

import "fmt"

type User struct {
	Id   int    `json:id`
	Name string `json:name`
}
