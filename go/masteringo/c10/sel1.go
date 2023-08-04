package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomGenerator(createRandom chan<- int, end <-chan bool) {
	rand.Seed(time.Now().UnixNano())
	for {
		select {
		case createRandom <- rand.Int():
			fmt.Println("Random number created!")
		case <-end:
			fmt.Println("Stop random generation")
			return
		case <-time.After(time.Second * 2):
			fmt.Println("Timeout!")
		}
	}
}

func main() {
	random := make(chan int)
	end := make(chan bool)

	go randomGenerator(random, end)

	for i := 0; i < 10; i++ {
		fmt.Println("random:", <-random)
		time.Sleep(time.Second / 2)
	}

	time.Sleep(time.Second * 3)
	end <- true
	time.Sleep(time.Second)
}
