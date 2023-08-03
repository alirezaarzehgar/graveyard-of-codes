package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Person struct {
	ID        int
	Name      string
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	ID    int
	Name  string   `gorm:"unique"`
	Users []Person `gorm:"many2many:user_languages;"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("l1.db"))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(&Person{}) {
		db.Migrator().DropTable(&Language{}, &Person{}, "user_languages")
	}
	db.AutoMigrate(&Language{}, &Person{})

	langs := []Language{
		{Name: "Persian"},
		{Name: "English"},
		{Name: "Spanish"},
		{Name: "Arabic"},
	}
	db.Create(&langs)
	db.Create([]*Person{
		{Name: "Ali", Languages: []Language{langs[0], langs[1]}},
		{Name: "Reza", Languages: []Language{langs[1], langs[0]}},
		{Name: "Mohmmad", Languages: []Language{langs[0], langs[3]}},
	})

	var u Person
	db.Preload("Languages").Take(&u)
	fmt.Println(u)

	var l Language
	db.Model(&Language{}).Preload("Users").Find(&l)
	fmt.Println(l)
}
