package main

import "fmt"

type MyI int

func (m *MyI) add(n int) {
	*m += MyI(n)
}

func (m *MyI) sub(n int) {
	*m -= MyI(n)
}

func (m *MyI) mul(n int) {
	*m *= MyI(n)
}

func (m *MyI) div(n int) {
	*m /= MyI(n)
}

func main() {
	var c MyI
	var i *MyI = &c
	i.add(2)
	fmt.Println(*i)

	var v **MyI
	v = &i
	(*v).mul(2)
	fmt.Println(**v)

}
