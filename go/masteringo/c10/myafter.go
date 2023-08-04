package main

import (
	"fmt"
	"time"
)

func main() {
	alert := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			alert <- true
			time.Sleep(time.Second)
		}
	}()

	for {
		select {
		case <-alert:
			fmt.Println("Hi!")
		case <-time.NewTimer(time.Second * 2).C:
			fmt.Println("timout!")
		}
	}
}
