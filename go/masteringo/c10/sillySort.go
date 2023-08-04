package main

import (
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for _, arg := range os.Args[1:] {
		n, err := strconv.Atoi(arg)
		if err != nil || n < 0 {
			print(". ")
		}

		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			time.Sleep(time.Duration(n) * time.Millisecond)
			print(n, " ")
		}(n)
	}

	wg.Wait()
	println()
}
