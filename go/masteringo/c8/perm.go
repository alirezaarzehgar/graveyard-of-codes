package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Parse()
	for _, filename := range flag.Args() {
		info, err := os.Stat(filename)
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println("Permission:", info.Mode().String())
	}
}
