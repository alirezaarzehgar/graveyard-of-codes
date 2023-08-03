package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/hints"
)

type User struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"not null"`
	Age    uint   `gorm:"check:age > 1"`
	Gender string `gorm:"check:gender in ('male', 'female'); not null"`
}

type UserAPI struct {
	ID   uint
	Name string
}

func Condition1(db *gorm.DB) *gorm.DB {
	return db.Not("name = ?", "U1")
}

func Condition2(db *gorm.DB) *gorm.DB {
	return db.Where("age > ?", 18)
}

func main() {
	db, err := gorm.Open(sqlite.Open("l1.db"))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(&User{}) {
		db.Migrator().DropTable(&User{})
	}
	db.AutoMigrate(&User{})

	var u User
	db.Where(&User{Gender: "male"}).Attrs(User{Age: 19}).FirstOrInit(&u, &User{Name: "Mohammad"})
	fmt.Println(u)

	for i := 1; i <= 150; i++ {
		gender := "male"
		if i%2 == 0 {
			gender = "female"
		}

		db.Create(&User{
			Name:   fmt.Sprintf("U%d", i),
			Age:    uint(10 + i%10),
			Gender: gender,
		})
	}

	var apis []UserAPI
	db.Model(&User{}).Limit(3).Find(&apis)
	fmt.Println(apis)

	var us []User
	db.Clauses(clause.Locking{Strength: "UPDATE"}).Limit(3).Find(&us)
	fmt.Println(us)

	db.Clauses(hints.New("hint")).Find(&u)
	fmt.Println(u)

	rows, err := db.Model(&User{}).Where("age = ?", 19).Rows()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var u User
		db.ScanRows(rows, &u)
		fmt.Println(u)
	}

	var names []string
	db.Model(&User{}).Offset(5).Limit(3).Pluck("age", &names)
	fmt.Println(names)

	var c int64
	db.Scopes(Condition1, Condition2).Find(&us).Count(&c)
	fmt.Println(c, us)
}
