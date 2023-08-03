package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const COUNT = 200000

func main() {
	var sum int32
	var wg sync.WaitGroup
	wg.Add(COUNT)

	for i := 0; i < COUNT; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt32(&sum, 1)
		}()
	}
	wg.Wait()
	fmt.Println("i:", sum)
}
