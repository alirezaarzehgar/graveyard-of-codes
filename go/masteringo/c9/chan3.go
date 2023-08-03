package main

import "fmt"

func main() {
	c := make(chan int, 10)

	for i := 0; i < 10; i++ {
		c <- i
	}
	for i := 0; i < 10; i++ {
		fmt.Println("i:", <-c)
	}

	close(c)
	read, ok := <-c
	fmt.Println(read, ok)
}
