package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "admin:admin@tcp(localhost:9010)/") //establish connection to MySQL database
	if err != nil {
		panic(err) // built in function in Go. Stop executing and print error message to the console
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
