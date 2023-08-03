package main

import "fmt"

func (v int) print() {
	fmt.Println(v)
}

func main() {
	s := 2
	s.print()
}
