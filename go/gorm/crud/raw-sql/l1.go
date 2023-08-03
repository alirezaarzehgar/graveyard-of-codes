package main

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("l1.db"))
	if err != nil {
		log.Fatal(err)
	}

	var us []User
	err = db.Raw("SELECT * FROM users LIMIT @name OFFSET @off", map[string]any{"name": 5}, sql.Named("off", 6)).Scan(&us).Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(us)

	db.Raw("CREATE TABLE test (`id` INTEGER PRIMARY KEY AUTOINCREMENT, `name` STRING NOT NULL)")
	db.Raw("DROP TABLE test")

	var u User
	stmt := db.Session(&gorm.Session{DryRun: true}).First(&u).Statement
	fmt.Println(stmt.SQL.String())
	fmt.Println(stmt.Vars)

	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(&User{}).Where("id > ?", 3).Limit(2).Offset(4).Find(&[]User{})
	})
	fmt.Println(sql)
}
