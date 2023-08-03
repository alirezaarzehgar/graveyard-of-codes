package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func printFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scnr := bufio.NewScanner(f)
	for scnr.Scan() {
		io.WriteString(os.Stdout, scnr.Text())
		io.WriteString(os.Stdout, "\n")
	}
	return nil
}

func main() {
	args := os.Args
	if len(args) == 1 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	for _, filename := range args[1:] {
		if err := printFile(filename); err != nil {
			log.Println(err)
		}
	}
}
