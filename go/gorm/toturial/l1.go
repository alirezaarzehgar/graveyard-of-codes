package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type U1 struct {
	ID   int
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("l1.db"))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(&U1{}) {
		db.Migrator().DropTable(&U1{})
	}
	db.AutoMigrate(&U1{})

	tx := db.Begin()
	tx.Create([]*U1{{Name: "Ali"}, {Name: "Mohammad"}, {Name: "Reza"}})
	tx.SavePoint("r1")
	tx.Create([]*U1{{Name: "Amir"}, {Name: "Gholamreza"}, {Name: "Fatemeh"}})
	tx.SavePoint("r2")
	tx.Create([]*U1{{Name: "Javad"}, {Name: "Sadra"}, {Name: "Saleh"}})
	tx.RollbackTo("r1")
	tx.Commit()

	var us []U1
	db.Find(&us)
	for _, u := range us {
		fmt.Println(u)
	}
}
