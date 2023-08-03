package controller

import (
	"github.com/alirezaarzehgar/hi/database"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	godotenv.Load()
	db = database.Init()
}
