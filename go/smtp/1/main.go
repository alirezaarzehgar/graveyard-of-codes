package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	mConf := getMailInfo()
	auth := authMail(mConf)
	err := sendBetterMail(mConf, auth, "Better test", "Hi poor people!", []string{mConf.mail})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sent successfully")
}
