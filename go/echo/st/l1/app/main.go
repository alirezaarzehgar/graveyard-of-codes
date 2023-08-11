package main

import (
	"example/db"
	"example/routes"
	"log"
	"os"
)

func main() {
	if err := db.Init(); err != nil {
		log.Fatal(err)
	}
	r := routes.Init()
	r.Logger.Fatal(r.Start(os.Getenv("RUNNING_ADDRESS")))
}
