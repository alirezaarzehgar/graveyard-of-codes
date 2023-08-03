package main

import (
	"fmt"
	"log"
)

func main() {
	var myInt any = 123

	i := myInt.(int)
	fmt.Println(i)
	j, ok := myInt.(bool)
	if !ok {
		log.Fatal("bad type for myIny")
	}
	fmt.Println(j)
}
