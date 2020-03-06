package models

import "fmt"

type User struct {
	Id   int    `json:id`
	Name string `json:name`
}

func (user *User) Save() {
	fmt.Println(user)
	db.Create(user)
}
