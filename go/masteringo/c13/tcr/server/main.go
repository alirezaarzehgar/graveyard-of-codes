package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(os.Args[0], "[address]")
		os.Exit(1)
	}

	server, err := net.Listen("tcp", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	for {
		con
	}
}
