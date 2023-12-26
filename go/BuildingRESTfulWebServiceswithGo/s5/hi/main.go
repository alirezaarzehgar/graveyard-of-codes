package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
	Email    string `bson:"email"`
}

func main() {
	s, err := mgo.Dial("127.0.0.1")
	if err != nil {
		log.Fatal("can't dial mongodb: ", err)
	}
	defer s.Close()

	c := s.DB("appdb").C("users")

	err = c.Insert(User{"ali", "123", "user@example.com"})
	if err != nil {
		log.Fatal("cant insert to appdb.users: ", err)
	}

	var u User
	if err = c.Find(bson.M{"username": "ali"}).One(&u); err != nil {
		log.Fatal("can't find record with ali name: ", err)
	}

	fmt.Println("record: ", u)
}
