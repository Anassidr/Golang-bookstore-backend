package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func Connect() {
	connectionString := "host=localhost port=5432 user=postgres dbname=go-bookstore sslmode=disable"
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic(err) // built in function in Go. Stop executing and print error message to the console
		return
	}

	defer db.Close()
}

func GetDB() *gorm.DB {
	return db
}
