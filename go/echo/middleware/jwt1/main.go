package main

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var secret = []byte("secret")

func Login(c echo.Context) error {
	jwt.
	return nil
}

func Home(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return errors.New("JWT token missing or invalid")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("failed to cast claims to MapClaims")
	}
	return c.JSON(http.StatusOK, claims)
}

func main() {
	e := echo.New()
	e.POST("/login", Login)
	g := e.Group("/", echojwt.JWT(secret))
	g.GET("home", Home)

	e.Logger.Fatal(e.Start(":8000"))
}
