package main

import (
	"container/ring"
	"fmt"
)

var size = 10

func main() {
	r := ring.New(size + 1)

	for i := 1; i < r.Len()*2; i++ {
		r.Value = i * 2
		r = r.Next()
	}
	r.Value = 50

	sum := 0
	r.Do(func(a any) {
		if a == nil {
			fmt.Println(0)
			return
		}
		t := a.(int)
		fmt.Println(t)
		sum += t
	})
	fmt.Println("sum:", sum)
}
