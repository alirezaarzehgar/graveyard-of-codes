package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	Model
	Name    string
	Age     uint
	JobName string
}

type Sql struct {
	Sql string
}

func (s Sql) TableName() string {
	return "sqlite_master"
}

func main() {
	db, err := gorm.Open(sqlite.Open("l2.db"))
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&User{})

	var sqls []Sql
	db.Find(&sqls)

	for _, v := range sqls {
		fmt.Println(v)
	}
}
