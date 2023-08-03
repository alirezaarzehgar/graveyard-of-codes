package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Job struct {
	Name        string
	Description string
}

type User struct {
	ID   uint   `gorm:"primaryKey;autoIncrement;autoIncrementIncrement:3"`
	Name string `gorm:"unique;not null;column:fuckoff"`
	Age  uint   `gorm:"not null;default:10;check:age > 1"`
	Job  Job    `gorm:"embedded;embeddedPrefix:job_;index"`
}

type Sql struct {
	Sql string
}

func (s Sql) TableName() string {
	return "sqlite_master"
}

func main() {
	db, err := gorm.Open(sqlite.Open("l3.db"))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(User{}) {
		db.Migrator().DropTable(User{})
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 12; i++ {
		db.Create(&User{Name: fmt.Sprintf("Ali %d", i)})
	}

	var sqls []Sql
	db.Find(&sqls)
	for _, v := range sqls {
		fmt.Println(v)
	}

	var users []User
	db.Find(&users)
	for _, v := range users {
		fmt.Printf("%#v\n", v)
	}
}
