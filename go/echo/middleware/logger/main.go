package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var JwtKey = []byte("secret_key")
var Users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Login(c echo.Context) error {
	var cred Credentials
	err := json.NewDecoder(c.Request().Body).Decode(&cred)
	if err != nil {
		return err
	}

	expectedPassword, ok := Users[cred.Username]
	if !ok || expectedPassword != cred.Password {
		return echo.ErrUnauthorized
	}

	expirationTime := time.Now().Add(time.Minute * 5)
	claim := &Claims{
		Username: cred.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return err
	}

	// c.SetCookie(&http.Cookie{
	// 	Name:    "token",
	// 	Value:   tokenString,
	// 	Expires: expirationTime,
	// })
	return c.JSON(http.StatusOK, map[string]any{
		"bearer": tokenString,
	})
}

func Home(c echo.Context) error {
	return nil
}

func Refresh(c echo.Context) error {
	return nil
}

func main() {
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		fmt.Println("This is a request!")
		return next
	})

	e.GET("/login", Login)
	g := e.Group("/", middleware.JWT(JwtKey))
	g.GET("log", func(c echo.Context) error {
		return c.String(http.StatusOK, "logged")
	}, middleware.Logger())
	g.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	g.GET("home", Home)
	g.GET("refresh", Refresh)

	e.Logger.Fatal(e.Start(":8000"))
}
