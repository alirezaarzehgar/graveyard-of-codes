package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// HasMany/One to many
type Post struct {
	ID       int
	Title    string
	Content  string
	Comments []Comment
}

// BelongsTo/One to one
type Comment struct {
	ID      int
	Content string
	PostID  int
	UserID  int
	User    User
}

type User struct {
	ID   int
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("l1.db"))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(&Post{}) {
		db.Migrator().DropTable(&Post{}, &Comment{}, &User{})
	}
	db.AutoMigrate(&User{}, &Comment{}, &Post{})

	db.Create(&Post{
		Title:   "Fist Post",
		Content: "Content",
		Comments: []Comment{
			{Content: "Hi", User: User{Name: "Ali"}},
			{Content: "Hello", User: User{Name: "Mohammad"}},
		},
	})

	var post Post
	db.Preload("Comments.User").Take(&post)
	fmt.Printf("%#v\n", post)
}
