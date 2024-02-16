package main

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// dsn := "food_delivery:buonlamchiemoi147@tcp(127.0.0.1:3307)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := os.Getenv("CONNECTION_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	log.Println(db, err)
}
