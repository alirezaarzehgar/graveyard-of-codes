package main

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type U1 struct {
	Data int
}

type F struct {
	Data int
}

func main() {
	db, err := gorm.Open(sqlite.Open("l1.db"))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(&U1{}) {
		db.Migrator().DropTable(&F{}, &U1{})
	}
	db.AutoMigrate(&U1{}, &F{})

	for i := 0; i < 25; i++ {
		db.Create(&U1{Data: i * i})
	}

	var u U1
	err = db.First(&u).Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)

	err = db.Last(&u).Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)

	r := db.Take(&u)
	if r.Error != nil {
		log.Fatal(r.Error)
	}
	fmt.Println(u)
	fmt.Println("RowAffected: ", r.RowsAffected)

	var us []U1
	r = db.Find(&us)
	if r.Error != nil {
		log.Fatal(r.Error)
	}
	fmt.Println("RowAffected: ", r.RowsAffected)

	var ui map[string]any
	r = db.Model(&U1{}).Last(&ui)
	if r.Error != nil {
		log.Fatal(r.Error)
	}
	fmt.Printf("%#v\n", ui)

	var f F
	r = db.First(&f)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		log.Println("Record not found!", r.Error)
	}

	err = db.Find(&u, 12*12).Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)

	err = db.Find(&u, "data = ?", 2*2).Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
}
