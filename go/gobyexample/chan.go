package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	f := make(chan bool)

	go func() {
		time.Sleep(time.Second)
		f <- true
	}()

	go func() {
		<-f
		ch <- "Hello world"
	}()

	fmt.Println(<-ch)
}
