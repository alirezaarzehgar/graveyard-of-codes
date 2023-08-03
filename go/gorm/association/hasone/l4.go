package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// type Comment struct {
// 	ID          uint
// 	Title       string
// 	Description string
// 	OwnerID     uint
// 	OwnerType   string
// }

// type Video struct {
// 	ID       uint
// 	Title    string
// 	FilePath string
// 	Comment  Comment `gorm:"polymorphic:Owner"`
// }

// type Post struct {
// 	ID      uint
// 	Title   string
// 	Content string
// 	Comment Comment `gorm:"polymorphic:Owner"`
// }

type Toy struct {
	ID        int
	Name      string
	OwnerID   int
	OwnerType string
}

type Cat struct {
	ID   int
	Name string
	Toy  Toy `gorm:"polymorphic:Owner"`
}

type Dog struct {
	ID   int
	Name string
	Toy  Toy `gorm:"polymorphic:Owner"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("l4.db"))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(&Toy{}) {
		db.Migrator().DropTable(&Dog{}, &Cat{}, &Toy{})
	}
	db.AutoMigrate(&Toy{}, &Cat{}, &Dog{})
	db.Create(&Dog{Name: "dog1", Toy: Toy{Name: "toy1"}})
	db.Create(&Cat{Name: "cat1", Toy: Toy{Name: "toy2"}})

	var dogs []Dog
	db.Preload("Toy").Find(&dogs)
	fmt.Println(dogs)

	var cats []Cat
	db.Preload("Toy").Find(&cats)
	fmt.Println(cats)
}
