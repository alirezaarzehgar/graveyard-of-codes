package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Users struct {
	FirstName string
	LastName  string
	Age       uint
}

type SqliteMaster struct {
	Sql string
}

func (s SqliteMaster) TableName() string {
	return "sqlite_master"
}

func main() {
	db, err := gorm.Open(sqlite.Open("l1.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(Users{}) {
		db.Migrator().DropTable(&Users{})
	}
	db.AutoMigrate(&Users{})

	var sqls []SqliteMaster
	db.Find(&sqls)

	for _, v := range sqls {
		fmt.Println(v.Sql)
	}
}
