package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ID:        "1",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
	})
	strToken, _ := token.SignedString([]byte("secret"))
	fmt.Println(strToken)

	clams, ok := token.Claims.(jwt.RegisteredClaims)
	if !ok {
		log.Fatal("He'y")
	}
	fmt.Println(clams)
}
