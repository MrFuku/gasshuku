package models

import (
	"fmt"
	"os"

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
	CONNECT := fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DB_HOST"),
		os.Getenv("MYSQL_DB"),
	)
	db, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err)
	}
}
