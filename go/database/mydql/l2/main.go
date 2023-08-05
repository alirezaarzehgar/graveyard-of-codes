package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name string
}

func main() {
	dsn := "ali:123@tcp(127.0.0.1:3306)/test_hi"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(&Product{}) {
		db.Migrator().DropTable(&Product{})
	}
	db.AutoMigrate(&Product{})
}
