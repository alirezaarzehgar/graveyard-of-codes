package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var signal = make(chan struct{})
var DATA = make(map[int]bool)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func first(min, max int, out chan<- int) {
	for {
		select {
		case <-signal:
			close(out)
			return
		case out <- random(min, max):
		}
	}
}

func second(out chan<- int, in <-chan int) {
	for x := range in {
		fmt.Print(x, " ")
		_, ok := DATA[x]
		if ok {
			signal <- struct{}{}
		} else {
			DATA[x] = true
			out <- x
		}
	}
	fmt.Println()
	close(out)
}

func third(in <-chan int) {
	var sum int
	for x := range in {
		sum += x
	}
	fmt.Println("The sum of the random numbers is ", sum)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Need two integer parameters!")
		return
	}
	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])
	if n1 > n2 {
		fmt.Printf("%d should be smaller than %d\n", n1, n2)
		return
	}

	rand.Seed(time.Now().UnixNano())
	A := make(chan int)
	B := make(chan int)

	go first(n1, n2, A)
	go second(B, A)
	third(B)
}
