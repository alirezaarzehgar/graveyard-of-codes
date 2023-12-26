package main

import (
	"context"
	"log"
	"rest/app"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	client, err := mongo.Connect(app.CTX, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("appdb")

	ucol := db.Collection("users")
	_, err = ucol.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.M{"username": 1},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		log.Fatal(err)
	}

	user := app.UserHandler{Coll: ucol}
	e := echo.New()
	e.GET("/user/:id", user.GetUser)
	e.GET("/user/all", user.GetUsers)
	e.POST("/user/new", user.CreateUser)
	e.PUT("/user/:id", user.UpdateUser)
	e.DELETE("/user/:id", user.DeleteUser)
	e.Logger.Fatal(e.Start(":8000"))
}
