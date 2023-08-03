package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)

	go func(reciever <-chan string) {
		defer wg.Done()
		fmt.Println(<-reciever)
	}(c)
	go func(sender chan<- string) {
		defer wg.Done()
		sender <- "Hello, World"
	}(c)
	wg.Wait()
}
