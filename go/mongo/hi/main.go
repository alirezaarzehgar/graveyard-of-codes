package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
	City string `bson:"city"`
}

func main() {
	co := options.Client().ApplyURI(os.Getenv("CONN_STR"))
	c, err := mongo.Connect(context.TODO(), co)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := c.Disconnect(context.TODO()); err != nil {
			log.Fatal("disconnect:", err)
		}
	}()

	if err := c.Ping(context.TODO(), nil); err != nil {
		log.Print("ping:", err)
	}
	fmt.Println("Connected to MongoDB")

	collection := c.Database("test").Collection("trainers")

	r, err := collection.InsertOne(context.TODO(), Trainer{Name: "hamid", Age: 14, City: "Mashhad"})
	if err != nil {
		log.Fatal("insert one:", err)
	}

	fmt.Println("Insert a single document:", r.InsertedID)

	var rt Trainer
	err = collection.FindOne(context.TODO(), bson.D{{Key: "name", Value: "hamid"}}).Decode(&rt)
	if err != nil {
		log.Fatal("find one:", err)
	}

	fmt.Println(rt)
}
