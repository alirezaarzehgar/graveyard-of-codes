package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"<-:create;privateKey"`
	Name     string `gorm:"<-;unique;not null"`
	Age      uint   `gorm:"<-:create;->;check:age > 1;not null"`
	Password string `gorm:"<-;->:false; not null"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("l4.db"))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(&User{}) {
		db.Migrator().DropTable(User{})
	}
	db.AutoMigrate(&User{})

	db.Create(&User{Name: "Alireza", Age: 19, Password: "1234"})
	var u User
	db.Last(&u)
	fmt.Printf("%#v\n", u)

	u.Age = 23
	db.Save(&u)
	u = User{}
	db.Last(&u)
	fmt.Printf("%#v\n", u)

}
