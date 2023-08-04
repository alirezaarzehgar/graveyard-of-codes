package main

import (
	"fmt"
	"time"
)

func A(a, b chan struct{}) {
	<-a
	fmt.Println("A()!")
	time.Sleep(time.Second)
	close(b)
}

func B(a, b chan struct{}) {
	<-a
	fmt.Println("B()!")
	close(b)
}

func C(a chan struct{}) {
	<-a
	fmt.Println("C()!")
}

func main() {
	c1 := make(chan struct{})
	c2 := make(chan struct{})
	c3 := make(chan struct{})

	go C(c3)
	go A(c1, c2)
	go C(c3)
	go B(c2, c3)
	go C(c3)

	close(c1)
	time.Sleep(time.Second * 3)
}
