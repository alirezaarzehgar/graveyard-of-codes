package app

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var CTX = context.Background()

type User struct {
	ID       string `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type UserHandler struct {
	Coll *mongo.Collection
}

func (u *UserHandler) GetUser(c echo.Context) error {
	var user User
	err := u.Coll.FindOne(CTX, bson.M{"_id": c.Param("id")}).Decode(&user)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func (u *UserHandler) GetUsers(c echo.Context) error {
	var users []User
	cur, err := u.Coll.Find(CTX, bson.M{})
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	for cur.Next(CTX) {
		var user User
		cur.Decode(&user)
		users = append(users, user)
	}

	return c.JSON(http.StatusOK, users)
}

func (u *UserHandler) CreateUser(c echo.Context) error {
	var user User
	if err := json.NewDecoder(c.Request().Body).Decode(&user); err != nil {
		return echo.ErrBadRequest
	}

	user.ID = bson.NewObjectId().Hex()
	ur, err := u.Coll.InsertOne(CTX, user)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, ur)
}

func (u *UserHandler) UpdateUser(c echo.Context) error {
	var user User
	if err := json.NewDecoder(c.Request().Body).Decode(&user); err != nil {
		return echo.ErrBadRequest
	}

	user.ID = c.Param("id")
	ur, err := u.Coll.UpdateByID(CTX, user.ID, bson.M{"$set": user})
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, ur)
}

func (u *UserHandler) DeleteUser(c echo.Context) error {
	ur, err := u.Coll.DeleteOne(CTX, bson.M{"_id": c.Param("id")})
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, ur)
}
