package main

import (
	"errors"
	"fmt"
	"log"
)

func mul2(a *int) (int, error) {
	if a == nil {
		return 0, errors.New("Invalid pointer")
	}

	*a *= 2
	return *a, nil
}

func main() {
	a := 3
	fmt.Println(a)

	val, err := mul2(&a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a, val)
}
