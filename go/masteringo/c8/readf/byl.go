package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func lineByLine(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("error reading file", err)
			break
		}
		fmt.Print(line)
	}

	return nil
}

func main() {
	flag.Parse()

	for _, filename := range flag.Args() {
		if err := lineByLine(filename); err != nil {
			log.Println(err)
		}
	}
}
