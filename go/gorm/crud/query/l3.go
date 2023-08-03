package main

import (
	"fmt"
	"log"
	"math/rand"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Customer struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
}

type Order struct {
	ID         uint `gorm:"primaryKey"`
	CustomerId uint `gorm:"unique"`
}

type JoinContainer struct {
	Customer
	Order
}

func main() {
	db, err := gorm.Open(sqlite.Open("l3.db"))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(&Customer{}) {
		db.Migrator().DropTable(&Customer{}, &Order{})
	}
	db.AutoMigrate(&Customer{}, &Order{})

	names := []string{"Ali", "Mohmmad", "Reza", "Hamid", "Ahmad", "Mehran", "Milad", "Mostafa"}
	dbb := db.Begin()
	for _, name := range names {
		dbb.Create(&Customer{Name: name})
	}
	dbb.Commit()

	orders := []uint{1, 2, 3, 4, 5, 6}
	for i := len(orders) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		orders[i], orders[j] = orders[j], orders[i]
	}

	dbb = db.Begin()
	for _, order := range orders {
		dbb.Create(&Order{CustomerId: order})
	}
	dbb.Commit()

	var r []JoinContainer
	db.Table("customers").Select("*").Joins("inner join orders on customers.id = orders.customer_id").Scan(&r)
	fmt.Println(r)

	db.Model(&Customer{}).Select("customers.name, orders.customer_id").
		Joins("left join orders on customers.id = orders.customer_id").Scan(&r)
	fmt.Println(r)

	db.Model(&Order{}).Select("*").InnerJoins("customers").Scan(&r)
	fmt.Println(r)

	db.Model(&Customer{}).Select("*").InnerJoins("orders").Scan(&r)
	fmt.Println(r)

	var cs []Customer
	db.InnerJoins("customers").Find(&cs)
	fmt.Println(cs)

	var os []Order
	db.InnerJoins("orders").Find(&os)
	fmt.Println(os)
}
