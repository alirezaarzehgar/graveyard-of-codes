package main

import "fmt"

func main() {
	var a []int

	for i := 0; i < 100; i++ {
		a = append(a, i)
		fmt.Println(len(a), cap(a))
	}
}
