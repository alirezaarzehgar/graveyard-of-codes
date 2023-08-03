package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		select {
		case t := <-time.Tick(time.Second):
			fmt.Println(t)
		}
	}
}
