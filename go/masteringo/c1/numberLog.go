package main

import (
	"log"
	"os"
)

const LOGFILE = "/tmp/my.log"

func main() {
	f, err := os.OpenFile(LOGFILE, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	iLog := log.New(f, "myLogger", log.LstdFlags|log.Lshortfile)
	iLog.Print("hi hi")
	iLog.Print("hoy hoy")
}
