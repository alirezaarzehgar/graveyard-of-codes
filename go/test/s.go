package main

import "fmt"

func fact(x int) int {
	if x <= 1 {
		return x
	}
	return sum(x-1) * x
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(fact(i), ",")
	}
}
