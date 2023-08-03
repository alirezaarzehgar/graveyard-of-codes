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
	Team      []User `gorm:"foreignKey:ManagerID"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("l5.db"))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(&User{}) {
		db.Migrator().DropTable(&User{})
	}
	db.AutoMigrate(&User{})

	names := []string{"Ali", "Reza", "Hamid", "Amin", "Sadegh", "Hashem"}

	dbb := db.Begin()
	for _, name := range names {
		dbb.Create(&User{Name: name})
	}
	dbb.Commit()

	var users []User
	db.Find(&users)
	for _, user := range users {
		db.Not("name = ?", user.Name).Order("RANDOM()").Limit(3).Find(&user.Team)
		db.Save(&user)
	}

	users = nil
	db.Preload("Team").Find(&users)
	for _, user := range users {
		fmt.Println(user.Name, ":", user.Team)
	}
}
