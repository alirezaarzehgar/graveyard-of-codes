package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Customer struct {
	ID       int
	Name     string
	Products []Product `gorm:"many2many:customer_products;"`
}

type Product struct {
	ID        int
	Name      string
	Customers []Customer `gorm:"many2many:customer_products;"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("l2.db"))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(&Customer{}) {
		db.Migrator().DropTable(&Customer{}, &Product{}, "customer_products")
	}
	db.AutoMigrate(&Customer{}, &Product{})

	products := []Product{
		{Name: "Pen"},
		{Name: "Book"},
		{Name: "Mobile"},
		{Name: "Nife"},
		{Name: "Paper"},
		{Name: "Headset"},
		{Name: "RubberDuck"},
		{Name: "Marker"},
		{Name: "Table"},
		{Name: "Order"},
	}
	db.Create(&products)
	db.Create([]Customer{
		{Name: "Ali", Products: products[:4]},
		{Name: "Mohammad", Products: products[2:6]},
		{Name: "Reza", Products: products[4:8]},
		{Name: "Hamid", Products: products[6:10]},
		{Name: "Ahamd", Products: products[:10]},
	})

	var prods []Product
	fmt.Println("Fetch products:")
	db.Preload("Customers").Find(&prods)
	for _, product := range prods {
		fmt.Println(product)
	}

	var custs []Customer
	fmt.Println("Fetch customers:")
	db.Preload("Products").Find(&custs)
	for _, cust := range custs {
		fmt.Println(cust)
	}
}
