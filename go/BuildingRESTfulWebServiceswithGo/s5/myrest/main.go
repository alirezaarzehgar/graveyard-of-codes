package main

import (
	"log"
	"myrest/api/handler"
	"myrest/api/route"
	"myrest/database"
)

//	@Title			mongo test
//	@Description	simple test for mongodb
//	@Version		2.0

func main() {
	db, err := database.Init()
	if err != nil {
		log.Fatal("db err: ", err)
	}

	handler.SetDb(db)

	r := route.Init()
	r.Logger.Fatal(r.Start(":8000"))
}
