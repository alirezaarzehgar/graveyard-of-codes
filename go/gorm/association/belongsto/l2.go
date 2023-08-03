package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type X1 struct {
	ID int
	Y1 Y1 `gorm:"foreignKey:ID"`
}

type Y1 struct {
	ID int
}

type X2 struct {
	ID int
	Y2 Y2
}

type Y2 struct {
	ID   int
	X2ID int
}

func main() {
	db, err := gorm.Open(sqlite.Open("l2.db"))
	if err != nil {
		log.Fatal(err)
	}

	tables := []any{&X1{}, &Y1{}, &X2{}, &Y2{}}
	if db.Migrator().HasTable(&X1{}) {
		db.Migrator().DropTable(tables...)
	}
	db.AutoMigrate(tables...)

	for i := 1; i <= 5; i++ {
		db.Create(&X1{Y1: Y1{ID: i}})
		db.Create(&X2{Y2: Y2{ID: i}})
	}

	var x1 X1
	db.Preload("Y1").Take(&x1)
	fmt.Println(x1)

	var x2 X2
	db.Preload("Y2").Take(&x2)
	fmt.Println(x2)
}
