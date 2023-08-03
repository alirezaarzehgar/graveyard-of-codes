package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i = 0
	for {
		time.Sleep(time.Second / 2)
		fmt.Println(i)
		i++

		if i == 4 {
			break
		}
	}

main_loop:
	for i := 1; i < 100; i++ {
		for j := 1; j < 100; j++ {
			if i%j == 2 {
				fmt.Println("Condition accepted: i%j =", i%j, "; i =", i, ", j =", j)
				continue main_loop
			}
		}
	}
}
