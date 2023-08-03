package main

import (
	"fmt"
	"sync"
	"time"
)

func hello() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello, World", i)
		time.Sleep(time.Millisecond * 1500 / 2)
	}
	fmt.Println("Finish Hoyya")
}

func bye() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bye bye, World", i)
		time.Sleep(time.Millisecond * 980 / 2)
	}
	fmt.Println("Finish Boyya")
}

func main1() {
	end := make(chan bool)

	go func() {
		hello()
		end <- true
	}()

	go func() {
		bye()
		end <- true
	}()

	<-end
	<-end
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		hello()
	}()

	go func() {
		defer wg.Done()
		bye()
	}()

	wg.Wait()
}
