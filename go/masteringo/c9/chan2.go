package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(c chan int, x int) {
		defer wg.Done()
		fmt.Println("1", x)
		c <- x
		close(c)
		fmt.Println("2", x)
	}(c, 12)
	time.Sleep(time.Second)
	fmt.Println("channel value: ", <-c)
	wg.Wait()

	if _, ok := <-c; ok {
		fmt.Println("Channel is opened")
	} else {
		fmt.Println("Channel is closed")
	}
}
