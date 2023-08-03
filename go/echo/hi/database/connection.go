package database

import (
	"fmt"
	"log"
	"os"

	"github.com/alirezaarzehgar/hi/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(os.Getenv("DB_PATH")))
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(model.User{}, &model.Post{})
	fmt.Println("migrate", os.Getenv("DB_PATH"))

	return db
}
