package main

import (
	"fmt"
	"math/rand"
	"time"
)

func add(c chan int) {
	sum := 0
	t := time.NewTimer(time.Second)

	for {
		select {
		case num := <-c:
			sum += num
		case <-t.C:
			c = nil
			fmt.Println("Sum", sum)
		}
	}
}

func send(c chan int) {
	for {
		c <- rand.Intn(10)
	}
}

func main() {
	var nc chan int
	select {
	case <-nc:
		fmt.Println("nil channel wont run")
	case <-time.After(1):
		fmt.Println("Run")
	}

	c := make(chan int)
	go add(c)
	go send(c)

	time.Sleep(time.Second * 3)
}
