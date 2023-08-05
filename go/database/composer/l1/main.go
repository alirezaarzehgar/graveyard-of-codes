package main

import (
	"log"

	"github.com/BaseMax/JokeGoServiceAPI/internal/storage"
)

func main() {
	store, err := storage.NewMysqlStorage(storage.MysqlConfig{
		Username: "user",
		Password: "password",
		DbName:   "db",
		Port:     3306,
		Host:     "localhost",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(store)
}
