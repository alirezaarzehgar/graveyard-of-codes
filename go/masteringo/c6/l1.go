package main

import "fmt"

func main() {
	fmt.Println(
		func(x int) (d int, s int) {
			d = x * 2
			s = x * x
			return
		}(5),
	)
}
