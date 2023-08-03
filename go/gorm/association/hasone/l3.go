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
	db, err := gorm.Open(sqlite.Open("name.db"))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(&User{}) {
		db.Migrator().DropTable(&User{})
	}
	db.AutoMigrate(&User{})

	names := []string{"Ali", "Reza", "Hamid", "Amin", "Sadegh", "Hashem"}

	var u *User = &User{}
	up := u
	for _, name := range names {
		u.Manager = &User{Name: name}
		u = u.Manager
	}
	db.Create(up)

	u = &User{}
	for u != nil {
		db.Preload("Manager").Find(&u, len(names))
		fmt.Println(*u)
		u = u.Manager
	}
}
