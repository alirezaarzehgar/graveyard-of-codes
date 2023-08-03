package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type People struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&People{})

	db.Create(&People{Name: "Alireza", Age: 19})

	var p People
	db.First(&p, "age = ?", 19)
	db.Model(p).Update("age", 22)
}
