package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	n := flag.Uint("n", 0, "Number of goroutines")
	flag.Parse()

	if len(flag.Args()) != 2 {
		fmt.Println("Please enter goroutine number")
	}

	var wg sync.WaitGroup
	for i := 0; i < int(*n); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("Goroutine ", i)
		}(i)
	}

	fmt.Printf("waitgroup: %#v\n", wg)
	wg.Wait()
	fmt.Println("End process")
}
