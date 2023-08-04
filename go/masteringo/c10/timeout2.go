package main

import (
	"fmt"
	"sync"
	"time"
)

func timeout(w *sync.WaitGroup, t time.Duration) bool {
	tmp := make(chan int)
	go func() {
		defer close(tmp)
		time.Sleep(time.Second)

		w.Wait()
	}()
	select {
	case <-tmp:
		return false
	case <-time.After(t):
		return true
	}
}

func main() {
	var w sync.WaitGroup
	w.Add(1)

	if timeout(&w, time.Second/2) {
		fmt.Println("Timed out!")
	} else {
		fmt.Println("OK")
	}

	w.Done()
	if timeout(&w, time.Second*2) {
		fmt.Println("Timed out!")
	} else {
		fmt.Println("OK")
	}
}
