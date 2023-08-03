package main

import (
	"fmt"
)

type Str struct {
	text string
}

func (s Str) Text() string {
	return s.text
}

func (s *Str) Reverse() {
	reversed := ""

	for i := len(s.text) - 1; i >= 0; i-- {
		reversed += string(s.text[i])
	}

	s.text = reversed
}

func f() int {
	a := []int{}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Somebody giveme hoyya")
		}
		fmt.Println("defer")
	}()

	return a[2]
}

func main() {
	s := Str{"Hello, World"}
	s.Reverse()

	fmt.Println(s.Text())

	fmt.Println(f())
}
