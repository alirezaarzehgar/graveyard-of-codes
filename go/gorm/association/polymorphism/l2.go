package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Video struct {
	ID      int
	Url     string
	Comment []Comment `gorm:"polymorphic:Comment"`
}

type Picture struct {
	ID      int
	Url     string
	Comment []Comment `gorm:"polymorphic:Comment"`
}

type Comment struct {
	ID          int
	Content     string
	CommentID   int
	CommentType string
}

func main() {
	db, err := gorm.Open(sqlite.Open("l2.db"))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(&Comment{}) {
		db.Migrator().DropTable(&Comment{}, &Video{}, &Picture{})
	}
	db.AutoMigrate(&Comment{}, &Video{}, &Picture{})

	db.Create(&Video{Url: "path1/", Comment: []Comment{
		{Content: "hi"},
		{Content: "hello"},
		{Content: "bye"},
	}})
	db.Create(&Picture{Url: "path2/", Comment: []Comment{
		{Content: "how are u"},
		{Content: "thanks"},
		{Content: "whatsup"},
	}})

	var v Video
	db.Preload("Comment").Last(&v)
	fmt.Println(v)

	var p Picture
	db.Preload("Comment").Last(&p)
	fmt.Println(p)
}
