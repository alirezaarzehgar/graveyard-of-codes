package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

type Client struct {
	ID      int
	Integer int
}

type Data struct {
	Job    Client
	Square int
}

var (
	size    = 10
	data    = make(chan Data, size)
	clients = make(chan Client, size)
)

func worker(w *sync.WaitGroup) {
	for c := range clients {
		data <- Data{Job: c, Square: c.Integer * c.Integer}
		time.Sleep(time.Second / 2)
	}
	w.Done()
}

func main() {
	nJobPtr := flag.Int("job", 0, "number of jobs")
	nWorkerPtr := flag.Int("worker", 0, "number of workers")
	flag.Parse()

	go func() {
		for i := 0; i < *nJobPtr; i++ {
			clients <- Client{i, i}
		}
		close(clients)
	}()

	f := make(chan any)
	go func() {
		for d := range data {
			fmt.Print("Client ID:", d.Job.ID, "\tint:", d.Job.Integer)
			fmt.Print("\tsquare:", d.Square, "\n")
		}
		f <- true
	}()

	var w sync.WaitGroup
	for i := 0; i < *nWorkerPtr; i++ {
		w.Add(1)
		go worker(&w)
	}
	w.Wait()
	close(data)
	<-f
}
