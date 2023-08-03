package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type U2 struct {
	ID int `gorm:"primaryKey"`
	O1 int
	O2 int
	O3 int
	O4 int
	O5 int
}

func main() {
	db, err := gorm.Open(sqlite.Open("l2.db"))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(U2{}) {
		db.Migrator().DropTable(U2{})
	}
	db.AutoMigrate(&U2{})

	o := U2{0, 1, 2, 3, 4, 5}
	db.Select("o1", "o3").Create(&o)
	db.Omit("id", "o2", "o4").Create(&o)

	var us []U2
	for i := 0; i < 200; i++ {
		us = append(us, U2{0, 1, 2, 3, 4, 5})
	}

	_ = us
	db.CreateInBatches(&us, 10)

	us = nil
	db.Find(&us)
	for _, v := range us {
		fmt.Printf("%#v\n", v)
	}
}
