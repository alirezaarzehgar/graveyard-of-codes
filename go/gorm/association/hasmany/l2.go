package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Order struct {
	ID        int
	Customers []Customer
}

type Customer struct {
	ID      int
	Name    string
	OrderID int
}

func main() {
	db, err := gorm.Open(sqlite.Open("l2.db"))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(&Order{}) {
		db.Migrator().DropTable(&Customer{}, &Order{})
	}
	db.AutoMigrate(&Customer{}, &Order{})

	db.Create(&Order{Customers: []Customer{
		{Name: "Ali"}, {Name: "Reza"}, {Name: "Hamid"},
	}})

	var order Order
	db.Preload("Customers").Take(&order)
	fmt.Println(order)
}
