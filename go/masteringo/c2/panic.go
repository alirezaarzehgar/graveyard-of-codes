package main

import "fmt"

func main() {
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Message:", c)
		}
	}()
	panic("Fuck off")
}
