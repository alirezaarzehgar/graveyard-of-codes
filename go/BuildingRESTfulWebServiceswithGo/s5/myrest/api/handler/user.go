package handler

import (
	"encoding/json"
	"myrest/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2/bson"
)

// CreateUser godoc
//
//	@Summary		Create user to database
//	@Description	User creation
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		model.User	true	"user body"
//	@Success		200		{object}	model.User
//	@Failure		400		{object}	string
//
//	@Router			/user/new [POST]
func CreateUser(c echo.Context) error {
	var user model.User
	if err := json.NewDecoder(c.Request().Body).Decode(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"status": "failed", "err": err})
	}

	if err := users.Insert(user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"status": "failed", "err": err})
	}

	return c.JSON(http.StatusOK, map[string]any{"status": "ok"})
}

// GetUser godoc
//
//	@Summary	Get user by username
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Param		name	path		string	true	"username"
//	@Success	200		{object}	model.User
//	@Failure	400		{object}	string
//
//	@Router		/user/{name} [GET]
func GetUser(c echo.Context) error {
	var user model.User
	err := users.Find(bson.M{"username": c.Param("name")}).One(&user)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{"status": "failed", "err": "not found"})
	}
	return c.JSON(http.StatusOK, user)
}

// GetAllUsers godoc
//
//	@Summary		Get all users
//	@Description	fetch all users from db
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]model.User
//	@Failure		400	{object}	string
//
//	@Router			/user/all [GET]
func GetAllUsers(c echo.Context) error {
	var allUsers []model.User
	err := users.Find(bson.M{}).All(&allUsers)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{"status": "failed", "err": "not found"})
	}
	return c.JSON(http.StatusOK, allUsers)
}

// UpdateUser godoc
//
//	@Summary		update a user
//	@Description	update user by username
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			name	path		string		true	"username"
//	@Param			user	body		model.User	true	"user body"
//	@Success		200		{object}	model.User
//	@Failure		400		{object}	string
//
//	@Router			/user/{name} [PATCH]
func UpdateUser(c echo.Context) error {
	var user model.User
	if err := json.NewDecoder(c.Request().Body).Decode(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"status": "failed", "err": err})
	}

	users.Update(bson.M{"username": c.Param("name")}, user)
	return c.JSON(http.StatusOK, map[string]any{"status": "ok"})
}

// DeleteUser godoc
//
//	@Summary		delete a user
//	@Description	delete user by username
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			name	path		string	true	"username"
//	@Success		200		{object}	model.User
//	@Failure		400		{object}	string
//
//	@Router			/user/{name} [DELETE]
func DeleteUser(c echo.Context) error {
	if err := users.Remove(bson.M{"username": c.Param("name")}); err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{"status": "failed", "err": "not found"})
	}
	return c.JSON(http.StatusOK, map[string]any{"status": "deleted"})
}
