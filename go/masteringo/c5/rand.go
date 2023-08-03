package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	b := make([]byte, 50)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}

	str := base64.URLEncoding.EncodeToString(b)
	fmt.Println(str)
}
