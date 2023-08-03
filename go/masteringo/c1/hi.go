package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println("Hello, World")
	fmt.Fprintln(os.Stdout, "Hello, World")
	i, k := 3, 4
	j, k := 1, 2

	fmt.Println(i, j, k)
}
