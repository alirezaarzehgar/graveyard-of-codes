package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	var err error
	godotenv.Load(".env")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("MYSQL_DATABASE"))
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}
	return nil
}

func CreateConnection() *gorm.DB {
	if db == nil {
		Init()
	}
	return db
}
