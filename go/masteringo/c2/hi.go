package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("You'r using ", runtime.Compiler, " ")
	fmt.Println("on a", runtime.GOARCH, "machine")
	fmt.Println("Using go version", runtime.Version())
	fmt.Println("Number of CPU:", runtime.NumCPU())
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())

	fmt.Println("Hello, World!")
}
