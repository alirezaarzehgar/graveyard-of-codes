package main

import (
	"log"
	"log/syslog"
	"os"
	"path"
)

const LOGFILE = "/tmp/my.log"

func main() {
	programName := path.Base(os.Args[0])
	sysLog, err := syslog.New(syslog.LOG_ALERT, programName)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(sysLog)

	log.Println("somebody giveme hoyyya")
	log.Print("hi")

	iLog, err := os.OpenFile(LOGFILE, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer iLog.Close()

	log.SetOutput(iLog)

	log.Print("ey ghashang tar az paria")
	log.SetFlags(log.Ldate)
	log.Print("tanha to kooche naria")
}
