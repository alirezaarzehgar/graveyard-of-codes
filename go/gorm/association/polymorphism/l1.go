package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Image struct {
	ID            int
	Url           string
	ImageableID   int
	ImageableType string
}

type User struct {
	ID    int
	Name  string
	Image Image `gorm:"polymorphic:Imageable"`
}

type Post struct {
	ID    int
	Name  string
	Image Image `gorm:"polymorphic:Imageable"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("l1.db"))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(&Image{}) {
		db.Migrator().DropTable(&User{}, &Post{}, &Image{})
	}
	db.AutoMigrate(&User{}, &Post{}, &Image{})

	db.Create(&User{Name: "Ali", Image: Image{Url: "path1/"}})
	db.Create(&Post{Name: "Blah blah", Image: Image{Url: "path2/"}})

	var u User
	db.Preload("Image").Last(&u)
	fmt.Println(u)

	var p Post
	db.Preload("Image").Last(&p)
	fmt.Println(p)
}
