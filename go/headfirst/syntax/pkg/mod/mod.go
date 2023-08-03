package mod

import "fmt"

type Bar struct {
	Num  int
	Per  float64
	Name string
	tt   int
}

type Baz struct {
	Bar
	ff int
}

func (f *Baz) Hi() {
	fmt.Println("Hello, World")
	f.tt = 1
}
