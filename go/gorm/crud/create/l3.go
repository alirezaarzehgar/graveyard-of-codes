package main

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type U3 struct {
	Name string
	Age  uint
	Role string
}

func (u *U3) BeforeCreate(tx *gorm.DB) (err error) {
	if strings.ToLower(u.Role) == "admin" {
		return errors.New("Fuck off")
	}
	return
}

func (u *U3) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("Success", u.Name)
	return
}

func main() {
	db, err := gorm.Open(sqlite.Open("l3.db"))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(&U3{}) {
		db.Migrator().DropTable(&U3{})
	}
	db.AutoMigrate(&U3{})

	db.Create(&U3{Name: "Alireza", Age: 12, Role: "User"})
	err = db.Create(&U3{Name: "Alireza", Age: 12, Role: "Admin"}).Error
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			db.Session(&gorm.Session{SkipHooks: true}).
				Create(&U3{Name: fmt.Sprint("I", i), Age: 12, Role: "User"})
		} else {
			db.Create(&U3{Name: fmt.Sprint("I", i), Age: 12, Role: "User"})
		}
	}

	var users []map[string]any
	for i := 0; i < 100; i++ {
		user := map[string]any{"Name": fmt.Sprint("A", i), "Age": 10 + i, "Role": "User"}
		users = append(users, user)
	}

	err = db.Model(&U3{}).Create(&users).Error
	if err != nil {
		log.Fatal(err)
	}
}
