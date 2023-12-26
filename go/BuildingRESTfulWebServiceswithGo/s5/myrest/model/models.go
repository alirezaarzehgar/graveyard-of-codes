package model

type User struct {
	Username string `json:"username" bson:"username" example:"user"`
	Password string `json:"password" bson:"password" example:"password"`
}
