package handler

import "gopkg.in/mgo.v2"

var db *mgo.Database
var users *mgo.Collection

func SetDb(externalDb *mgo.Database) {
	db = externalDb
	users = db.C("users")
}
