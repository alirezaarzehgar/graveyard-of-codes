package main

import (
	"log"
	"log/syslog"
	"os"
)

func main1() {
	iLog, err := syslog.New(syslog.LOG_ALERT, "My log")
	if err != nil {
		log.Fatal()
	}
	defer iLog.Close()

	log.SetOutput(iLog)
	log.SetFlags(log.LUTC)
	log.Println("hello, world")

	f, err := os.OpenFile("/tmp/my.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal()
	}
	defer f.Close()

	log.SetOutput(f)
	log.SetFlags(log.Llongfile | log.Ldate | log.LstdFlags)
	log.Println("hi")

	syslog.
}
