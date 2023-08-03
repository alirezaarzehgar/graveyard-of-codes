package main

import "fmt"

func main() {
	var name string = "Seyed"
	name += fmt.Sprintf(" alireza(%d)", 23)

	fmt.Println(name)
}
