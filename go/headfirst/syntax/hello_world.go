package main

import (
	"fmt"
	"reflect"
	"strings"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Hello, World")
	fmt.Println(strings.Title("hello"))
	fmt.Printf("rune: %v\n", 'x')

	var condition bool = true
	if condition {
		fmt.Println("condition was true")
	} else {
		fmt.Println("condition was false")
	}

	lst := []interface{}{12, "Hello, World", 3.23, 100000000000000000, true, 's'}

	for _, e := range lst {
		fmt.Println(reflect.TypeOf(e), ": ", e)
	}

	var a, b int

	a, b = 10, 30
	fmt.Println(a, b)

	a, b = b, a
	fmt.Println(a, b)

	var num int
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	fmt.Printf("%v\n", lst)
	fmt.Printf("%+v\n", lst)
	fmt.Printf("%T\n", lst)

	fmt.Printf("%t\n", lst)
	fmt.Printf("%b\n", 123)
	fmt.Printf("%c\n", 'x')
	fmt.Printf("%o\n", 10)
	fmt.Printf("0x%X\n", 0b1111)

	stops := [][2]int{
		{3, 14},
		{3, 19},
		{3, 5},
		{4, 5},
		{5, 6},
		{6, 7},
		{7, 9},
		{8, 12},
	}

	for i := 0; i < 10; i++ {
	outer:
		for j := 0; j < 30; j++ {
			for _, e := range stops {
				if e[0] == i && e[1] == j {
					fmt.Printf("\033[36mX\033[0m")
					continue outer
				}
			}

			time.Sleep(time.Millisecond * 5)
			syscall.Write(syscall.Stdout, []byte("*"))
		}
		fmt.Println()
	}
}
