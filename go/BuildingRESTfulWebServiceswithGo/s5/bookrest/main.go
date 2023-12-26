package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var ctx context.Context

type User struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password,omitempty"`
}

type DB struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func (db *DB) getUser(c echo.Context) error {
	var u User
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	fmt.Println(id)
	err := db.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&u)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	fmt.Println(u.ID)
	return c.JSON(http.StatusOK, u)
}

func (db *DB) getUsers(c echo.Context) error {
	var us []User
	cur, err := db.collection.Find(ctx, bson.M{})
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var usr User
		cur.Decode(&usr)
		us = append(us, usr)
	}

	return c.JSON(http.StatusOK, us)
}

func (db *DB) postUser(c echo.Context) error {
	var u User
	json.NewDecoder(c.Request().Body).Decode(&u)
	ur, err := db.collection.InsertOne(ctx, u)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, ur)
}

func (db *DB) putUser(c echo.Context) error {
	return nil
}

func (db *DB) deleteUser(c echo.Context) error {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	u, err := db.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, u)

}

func main() {
	ctx = context.Background()

	client, err := mongo.Connect(
		ctx,
		options.Client().
			ApplyURI("mongodb://127.0.0.1:27017").
			SetMonitor(&event.CommandMonitor{
				Started: func(ctx context.Context, cse *event.CommandStartedEvent) {
					log.Println(cse.Command)
				},
			}),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	database := client.Database("appdb")
	uc := database.Collection("users")

	uc.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: []string{"username"},
	})

	db := DB{client, uc}

	e := echo.New()
	e.GET("/user/:id", db.getUser)
	e.GET("/user/all", db.getUsers)
	e.POST("/user/new", db.postUser)
	e.PUT("/user/:id", db.putUser)
	e.DELETE("/user/:id", db.deleteUser)
	e.Logger.Fatal(e.Start(":8000"))
}
