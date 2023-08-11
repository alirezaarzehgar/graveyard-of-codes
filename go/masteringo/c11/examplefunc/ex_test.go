package examplefunc

import "fmt"

func ExampleAdd() {
	fmt.Println(Add(1, 2))
	fmt.Println(Add(10, -2))
	// Output:
	// 3
	// 8
}

func ExampleMultiply() {
	fmt.Println(Multiply(2, 2))
	fmt.Println(Multiply(6, 5))
	// Output:
	// 4
	// 30
}
