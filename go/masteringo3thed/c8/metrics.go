package main

import (
	"fmt"
	"runtime/metrics"
	"sync"
	"time"
)

func main() {
	mets := make([]metrics.Sample, 1)
	mets[0].Name = "/sched/goroutines:goroutines"

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Second)
		}()
		metrics.Read(mets)
		fmt.Println("Number of goroutines :", mets[0].Value.Uint64())
	}

	wg.Wait()
	metrics.Read(mets)
	fmt.Println("Number of goroutines :", mets[0].Value.Uint64())
}
