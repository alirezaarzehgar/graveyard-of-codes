// Incomplete many2many relationship
package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Vid struct {
	ID   int
	Url  string
	Tags []Tag `gorm:"many2many:all_tags"`
}

type Pos struct {
	ID   int
	Url  string
	Tags []Tag `gorm:"many2many:all_tags"`
}

type Tag struct {
	ID   int
	Name string
}

type Taggable struct {
	ID           int
	TagID        int
	TaggableID   int
	TaggableType string
}

func main() {
	db, err := gorm.Open(sqlite.Open("l3.db"))
	if err != nil {
		log.Fatal(err)
	}

	tables := []any{&Vid{}, &Tag{}}
	if db.Migrator().HasTable(&Tag{}) {
		db.Migrator().DropTable(tables...)
	}
	db.AutoMigrate(tables...)

	var tags []Tag
	for i := 'a'; i <= 'k'; i++ {
		tags = append(tags, Tag{Name: fmt.Sprintf("%c", i)})
	}

	db.Create(&tags)
}
