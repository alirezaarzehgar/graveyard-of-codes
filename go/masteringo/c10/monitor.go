package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const REPEAT = 16

var readValue = make(chan int)
var writeValue = make(chan int)

func monitor() {
	var value int
	for {
		select {
		case newVal := <-writeValue:
			value = newVal
			fmt.Print(value, " ")
		case readValue <- value:
			fmt.Print(" ")
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	go monitor()

	var w sync.WaitGroup
	for i := 0; i < REPEAT; i++ {
		w.Add(1)
		go func() {
			defer w.Done()
			writeValue <- rand.Intn(10 * REPEAT)
		}()
	}
	w.Wait()
	fmt.Println("\nLast value:", <-readValue)
}
