package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID   int    `gorm:"id"`
	Name string `gorm:"name"`
}

func main() {
	db, err := sql.Open("mysql", "ali:123@tcp(127.0.0.1:3306)/test_hi")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Query("DROP TABLE IF EXISTS users")
	_, err = db.Query("CREATE TABLE users (id int AUTO_INCREMENT NOT NULL, name varchar(255) NOT NULL, PRIMARY KEY(id))")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		_, err = db.Query("INSERT INTO users (NAME) VALUES (?)", fmt.Sprint("Ali ", i))
		if err != nil {
			log.Println(err)
		}
	}

	gdb, err := gorm.Open(mysql.Open("ali:123@tcp(127.0.0.1:3306)/test_hi"))
	if err != nil {
		log.Fatal("open:", err)
	}

	var users []User
	if err := gdb.Find(&users).Error; err != nil {
		log.Fatal("find:", err)
	}
	for _, u := range users {
		fmt.Println(u)
	}
}
