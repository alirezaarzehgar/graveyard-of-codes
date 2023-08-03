package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Country struct {
	ID     int
	Name   string
	Cities []City
}

type City struct {
	ID        int
	Name      string
	CountryID int
}

func main() {
	db, err := gorm.Open(sqlite.Open("l3.db"))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(&Country{}) {
		db.Migrator().DropTable(&City{}, &Country{})
	}
	db.AutoMigrate(&City{}, &Country{})

	db.Create(&Country{Name: "Iran", Cities: []City{
		{Name: "Mashhad"},
		{Name: "Tehran"},
		{Name: "Tabriz"},
		{Name: "Esfahan"},
		{Name: "Mazandaran"},
	}})

	var country Country
	db.Preload("Cities").Take(&country)
	fmt.Printf("%#v\n", country)
}
