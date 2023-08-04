package main

import "fmt"

func main() {
	numbers := make(chan int, 5)
	coutner := 10

	for i := 0; i < coutner; i++ {
		select {
		case numbers <- i:
		default:
			fmt.Println("Not enough space for", i)
		}
	}

	for i := 0; i < coutner+5; i++ {
		select {
		case num := <-numbers:
			fmt.Println("num:", num)
		default:
			fmt.Println("Nothing more to be done!")
			break
		}
	}
}
