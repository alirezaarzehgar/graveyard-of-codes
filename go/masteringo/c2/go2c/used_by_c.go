package main

import "C"
import "fmt"

//export PrintMessage
func PrintMessage() {
	fmt.Println("Hello, World!")
}

//export Multiply
func Multiply(a, b int) int {
	return a * b
}

func main() {
}
