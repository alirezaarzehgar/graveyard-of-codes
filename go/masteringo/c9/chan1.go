package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func(c chan int, x int) {
		fmt.Println("1", x)
		c <- x // nobody read this channel. this function never done
		close(c)
		fmt.Println("2", x)
	}(c, 5)
	time.Sleep(time.Second)
}
