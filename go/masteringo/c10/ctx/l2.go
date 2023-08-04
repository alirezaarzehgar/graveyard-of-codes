package main

import (
	"context"
	"fmt"
	"time"
)

func f() {
	c := context.Background()
	c, cancel := context.WithTimeout(c, time.Second)
	// c, cancel := context.WithDeadline(c, time.Now().Add(time.Second))
	// c, cancel := context.WithCancel(c)

	go func() {
		fmt.Println("start task")
		time.Sleep(time.Second * 3)
		fmt.Println("end task")
		cancel()
	}()

	select {
	case <-c.Done():
		fmt.Println("Done:", c.Err())
	case t := <-time.After(time.Second * 10):
		fmt.Println("after:", t, c.Err())
	}
}

func main() {
	f()
}
