package main

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.Name == "Admin" {
		return errors.New("you can't change name to Admin")
	}
	return
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

	for i := 0; i < 150; i++ {
		db.Create(&User{Name: fmt.Sprint("U", i)})
	}

	var u User
	db.First(&u)
	u.Name = "Hosein"
	db.Save(u)

	db.Save(&User{ID: 2, Name: "Ali"})

	db.Model(&User{}).Where("id < ?", 5).Where("id > ?", 2).Updates(map[string]any{"name": "changed"})

	rows, err := db.Model(&User{}).Limit(5).Rows()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var u User
		db.ScanRows(rows, &u)
		fmt.Println(u)
	}

	err = db.Save(&User{ID: 5, Name: "Admin"}).Error
	if err != nil {
		log.Fatal(err)
	}
}
