package main

import (
	"net/smtp"
	"os"
	"strings"
)

type MailConf struct {
	mail, pass, host, port, servAddr string
}

func getMailInfo() MailConf {
	host := os.Getenv("MAIL_HOST")
	port := os.Getenv("MAIL_PORT")
	return MailConf{
		mail: os.Getenv("EMAIL_ADDRESS"),
		pass: os.Getenv("EMAIL_PASSWORD"),
		host: host, port: port,
		servAddr: host + ":" + port,
	}
}

func authMail(c MailConf) smtp.Auth {
	return smtp.PlainAuth("", c.mail, c.pass, c.host)
}

func sendBetterMail(c MailConf, a smtp.Auth, subject, body string, to []string) error {
	msg := []byte("From: " + c.mail + "\r\n" +
		"To: " + strings.Join(to, ",") + "\r\n" +
		"Subject: " + subject + "\r\n\r\n" +
		body,
	)
	return smtp.SendMail(c.servAddr, a, c.mail, to, msg)
}
