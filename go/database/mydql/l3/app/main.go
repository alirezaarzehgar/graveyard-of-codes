package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	dsn := "root:pass@(127.0.0.1:3307)/test"
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})
}
