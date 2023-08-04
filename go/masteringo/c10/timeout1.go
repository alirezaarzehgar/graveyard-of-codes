package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	go func() {
		time.Sleep(time.Second * 3)
		c1 <- "c1 OK"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}

	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second)
		c2 <- "c2 OK"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 2):
		fmt.Println("timeout")
	}
}
