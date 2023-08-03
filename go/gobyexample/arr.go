package main

import (
	"errors"
	"fmt"
)

func main() {
	var ss []string = make([]string, 3)
	ss = append(ss, "hello")
	ss = append(ss, "my")
	ss = append(ss, "world")

	for _, s := range ss {
		fmt.Println(s)
	}

	var a []string
	fmt.Println(a == nil)

	var m map[any]any = map[any]any{
		"hi":               2,
		3:                  "ali",
		errors.New("yaya"): []int{1, 2, 3},
	}

	fmt.Println("----------")
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("----------")
	fmt.Println("len(m):", len(m))
	fmt.Println("len(ss):", len(ss))

	delete(m, 3)
}
