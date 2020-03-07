package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func InitDB() {
	connectDB()
	db.AutoMigrate(&User{}, &Comment{})
}

func connectDB() {
	var err error
	DBMS := "mysql"
	CONNECT := "user:password@(db)/gasshuku_database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err)
	}
}
