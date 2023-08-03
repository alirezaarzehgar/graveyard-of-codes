package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintln(os.Stderr, "Pass one or more float")
		os.Exit(1)
	}

	min := math.Inf(1)
	max := math.Inf(-1)
	for _, str := range os.Args[1:] {
		n, err := strconv.ParseFloat(str, 64)
		if err != nil {
			log.Fatal("Invalid floating point number: '", str, "`")
		}

		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)

	io.WriteString(os.Stdout, "Hello, World\n")
	w := io.Writer(os.Stdout)
	sw, _ := w.(io.StringWriter)
	sw.WriteString()
}
