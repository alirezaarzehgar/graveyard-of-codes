package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func wordByWord(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scnr := bufio.NewReader(f)
	for {
		line, err := scnr.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("error reading file:", err)
			return err
		}

		for _, c := range line {
			fmt.Printf("%c\n", c)
		}
	}
	return nil
}

func main() {
	flag.Parse()
	for _, filename := range flag.Args() {
		if err := wordByWord(filename); err != nil {
			log.Println(err)
		}
	}
}
