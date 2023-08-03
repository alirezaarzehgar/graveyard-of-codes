package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type U4 struct {
	Val string
}

func main() {
	db, err := gorm.Open(sqlite.Open("l4.db"))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(&U4{}) {
		db.Migrator().DropTable(&U4{})
	}
	db.AutoMigrate(&U4{})

	db.Model(&U4{}).Create([]map[string]any{
		{"Val": clause.Expr{SQL: "(SELECT sqlite_version())"}},
		{"Val": clause.Expr{SQL: "(SELECT sqlite_version())"}},
		{"Val": clause.Expr{SQL: "(SELECT sqlite_version())"}},
		{"Val": "(SELECT sqlite_version())"},
		{"Val": clause.Expr{SQL: "((SELECT ?))", Vars: []any{"mohammad"}}},
	})

	var us []U4
	db.Find(&us)
	for _, v := range us {
		fmt.Println(v)
	}
}
