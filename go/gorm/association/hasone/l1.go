package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type U1 struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	CreditCard CreditCard
}

type CreditCard struct {
	ID     int `gorm:"primaryKey"`
	Number string
	UserID int
}

func main() {
	db, err := gorm.Open(sqlite.Open("l1.db"))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(&U1{}) {
		db.Migrator().DropTable(&U1{}, &CreditCard{})
	}
	db.AutoMigrate(&CreditCard{}, &U1{})

	var u U1
	db.Create(&U1{
		Name:       "Alireza",
		CreditCard: CreditCard{Number: "3253"},
	})

	u = U1{}
	db.Model(&U1{}).Preload("CreditCard").Find(&u)
	fmt.Println(u)
}
