package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type U1 struct {
	gorm.Model
	Name     string `gorm:"unique;not null"`
	Age      uint   `gorm:"not null; check:age > 0"`
	FuckedUp bool   `gorm:"-"`
}

func (u U1) TableName() string {
	return "users"
}

type S1 struct {
	Sql string
}

func (s S1) TableName() string {
	return "sqlite_master"
}

func main() {
	db, err := gorm.Open(sqlite.Open("l1.db"))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(&U1{}) {
		db.Migrator().DropTable(&U1{})
	}
	db.AutoMigrate(&U1{})

	for i := 0; i < 12; i++ {
		u := U1{Name: fmt.Sprint("Ali ", i), Age: uint(10 + i)}
		db.Create(&u)
	}

	var users []U1
	for i := 0; i < 8; i++ {
		users = append(users, U1{
			Name: fmt.Sprint("Mohmmad ", i),
			Age:  uint(10 + i),
		})
	}
	db.Create(&users)

	var ss []S1
	db.Find(&ss)
	fmt.Println("Data Definition Commands:")
	for _, v := range ss {
		fmt.Println(v)
	}

	var us []U1
	db.Select("name", "age").Find(&us)
	fmt.Println("Data Manipulation Commands:")
	for _, v := range us {
		fmt.Println(v)
	}
}
