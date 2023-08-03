package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	io.WriteString(os.Stderr, "Hello world\n")

	s := strings.NewReader("Ba man sanama del del del del deldedledledleldeldledl yek dele koooon")
	s.WriteTo(os.Stderr)
	fmt.Println()

	b := make([]byte, 11)
	s.ReadAt(b, 22)
	fmt.Println(string(b))

	s.Seek(io.SeekStart, os.SEEK_SET)
	for {
		b = make([]byte, 5)
		n, err := s.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Printf("Read: %s Bytes: %d\n", b, n)
	}

}
