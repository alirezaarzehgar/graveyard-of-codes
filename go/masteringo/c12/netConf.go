package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range interfaces {
		fmt.Printf("%#v\n", i)
	}
}
