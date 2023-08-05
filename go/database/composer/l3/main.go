package main

import (
	"flag"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	h := flag.String("host", "db", "get hostname")
	flag.Parse()
	host := *h

	dsn := "user:pass@tcp(" + host + ":3306)/test"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}

	sqldb, _ := db.DB()
	if err = sqldb.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("You win bro!")
}
