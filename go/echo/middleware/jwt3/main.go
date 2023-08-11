package main

import (
	"hi/api"
	"hi/model"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	err := model.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	e := api.InitRoutes()
	e.Logger.Fatal(e.Start(":8000"))
}
