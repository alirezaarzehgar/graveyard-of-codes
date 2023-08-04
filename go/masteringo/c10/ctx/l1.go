package main

import (
	"context"
	"flag"
	"fmt"
	"time"
)

func f1(t int) {
	c1 := context.Background()
	c1, cancel := context.WithCancel(c1)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c1.Done():
		fmt.Println("f1():", c1.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f1():", r)
	}
	return
}
func f2(t int) {
	c2 := context.Background()
	c2, cancel := context.WithTimeout(c2, time.Second*time.Duration(t))
	defer cancel()

	go func() {
		time.Sleep(time.Second * 4)
		cancel()
	}()

	select {
	case <-c2.Done():
		fmt.Println("f2():", c2.Err())
		return
	case r := <-time.After(time.Second * time.Duration(t)):
		fmt.Println("f2():", r)
	}
	return
}

func f3(t int) {
	c3 := context.Background()
	deadline := time.Now().Add(time.Duration(2*t) * time.Second)
	c3, cancel := context.WithDeadline(c3, deadline)
	defer cancel()

	go func() {
		time.Sleep(time.Second * 4)
		cancel()
	}()

	select {
	case <-c3.Done():
		fmt.Println("f3():", c3.Err())
		return
	case r := <-time.After(time.Second * time.Duration(t)):
		fmt.Println("f3():", r)
	}
	return
}

func f(t int, methd func(t int, c1 context.Context) (context.Context, context.CancelFunc)) {
	c1 := context.Background()
	c1, cancel := methd(t, c1)
	defer cancel()

	go func() {
		time.Sleep(time.Second * 4)
		cancel()
	}()

	select {
	case <-c1.Done():
		fmt.Println("f3():", c1.Err())
		return
	case r := <-time.After(time.Second * time.Duration(t)):
		fmt.Println("f3():", r)
	}
	return
}

func main() {
	n := flag.Int("n", 0, "get number")
	flag.Parse()

	delay := *n
	// f(delay, func(t int, c1 context.Context) (context.Context, context.CancelFunc) {
	// 	return context.WithCancel(c1)
	// })
	// f(delay, func(t int, c1 context.Context) (context.Context, context.CancelFunc) {
	// 	return context.WithDeadline(c1, time.Now().Add(time.Second))
	// })
	// f(delay, func(t int, c1 context.Context) (context.Context, context.CancelFunc) {
	// 	return context.WithCancel(c1)
	// })

	fmt.Println("delay:", delay)
	f1(delay)
	f2(delay)
	f3(delay)
}
