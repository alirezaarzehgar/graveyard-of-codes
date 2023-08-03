package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID        uint
	Name      string
	ManagerID *uint
	Manager   *User
}

func main() {
	db, err := gorm.Open(sqlite.Open("l2.db"))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(&User{}) {
		db.Migrator().DropTable(&User{})
	}
	db.AutoMigrate(&User{})

	db.Create(
		&User{Name: "Alireza",
			Manager: &User{Name: "Hosein",
				Manager: &User{Name: "Reza"}}})

	u := &User{Name: "Ahmad"}
	db.Last(&u.Manager)
	db.Create(u)

	db.Preload("Manager.Manager.Manager").Find(u, 4)
	for u != nil {
		fmt.Println(*u)
		u = u.Manager
	}
}
