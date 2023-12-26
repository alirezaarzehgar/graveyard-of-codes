package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

func main1() {

	// Sender data.
	from := "alirezaarzehgar82@gmail.com"
	password := "???"

	// Receiver email address.
	to := []string{
		"alirezaarzehgar82@gmail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("From: " + from + "\r\n" +
		"To: " + strings.Join(to, ",") + "\r\n" +
		"Subject: Just for test\r\n\r\n" +
		"Email: Message!")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
