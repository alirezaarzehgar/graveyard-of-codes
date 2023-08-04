package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m  sync.Mutex
	v1 int
)

func change(i int) {
	m.Lock()
	time.Sleep(time.Second)
	v1++
	if v1%10 == 0 {
		v1 -= i * 10
	}
	m.Unlock()
}

func read() int {
	m.Lock()
	a := v1
	m.Unlock()
	return a
}

func main() {
	var wg sync.WaitGroup
	fmt.Println(read())
	for i := 0; i < 21; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			change(i)
			fmt.Print("->", read(), " ")
		}(i)
	}

	wg.Wait()
	fmt.Println("\n", read())
}
