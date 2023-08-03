package main

import (
	"fmt"
	"runtime"
)

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func profile() {
	var stat runtime.MemStats
	runtime.ReadMemStats(&stat)

	fmt.Println("stat.HeapObjects: ", stat.HeapObjects)
}

func main() {
	var clst [10000]func() int
	var ilst [10000]int

	profile()
	for i := 0; i < 10000; i++ {
		ilst[i] = 12
	}
	profile()
	for i := 0; i < 10000; i++ {
		clst[i] = intSeq()
	}
	profile()
}
