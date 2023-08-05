package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	user := os.Getenv("USER")
	pass := os.Getenv("PASS")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println("dsn:", dsn)

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}
	_ = db
}
