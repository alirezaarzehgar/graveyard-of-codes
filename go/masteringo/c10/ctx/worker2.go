package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

var (
	size   = 10
	data   = make(chan int, size)
	client = make(chan int, size)
)

func main() {
	nJobs := flag.Int("njobs", 10, "number of jobs")
	nWorkers := flag.Int("nworkers", 2, "number of workers")
	flag.Parse()

	go func() {
		for i := 0; i < *nJobs; i++ {
			client <- i
		}
		close(client)
	}()

	f := make(chan any)
	go func() {
		for d := range data {
			fmt.Println("Square:", d)
		}
		f <- true
	}()

	var w sync.WaitGroup
	for i := 0; i < *nWorkers; i++ {
		w.Add(1)
		go func() {
			for c := range client {
				data <- c * c
				time.Sleep(time.Second)
			}
			w.Done()
		}()
	}

	w.Wait()
	close(data)
	<-f
}
