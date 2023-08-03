package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID        int
	Name      string
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	ID    int
	Name  string
	Users []User `gorm:"many2many:user_languages;"`
}

func printData[T any](msg string, list []T) {
	fmt.Println(msg, ":")
	for _, item := range list {
		fmt.Println(item)
	}
	fmt.Println()
}

func main() {
	db, err := gorm.Open(sqlite.Open("l1.db"))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(&User{}) {
		db.Migrator().DropTable(&User{}, &Language{})
	}
	db.AutoMigrate(&User{}, &Language{})

	langs := []Language{
		{Name: "Persian"},
		{Name: "English"},
		{Name: "Spanish"},
		{Name: "Turkish"},
		{Name: "Arabic"},
	}
	db.Create(langs)
	db.Create([]User{
		{Name: "Ali", Languages: []Language{langs[0], langs[2]}},
		{Name: "Mohammad", Languages: []Language{langs[3], langs[4]}},
		{Name: "Reza", Languages: []Language{langs[0], langs[2]}},
		{Name: "Hosein", Languages: []Language{langs[1], langs[3]}},
		{Name: "Hamid", Languages: []Language{langs[3], langs[0]}},
		{Name: "Ahmad", Languages: []Language{langs[2], langs[4]}},
	})

	var users []User
	db.Preload("Languages").Find(&users)
	printData("User", users)

	var languages []Language
	db.Preload("Users").Find(&languages)
	printData("Languages", languages)

	db.Model(&users).Association("Languages").Delete(langs[:3])

	db.Preload("Languages").Find(&users)
	printData("User", users)

	db.Model(&users).Association("Languages").Clear()

	db.Preload("Users").Find(&languages)
	printData("Languages", languages)
}
