package main

import (
	"fmt"

	"github.com/alirezaarzehgar/hello/mod"
)

type Foo int

func (f *Foo) Add(x int) *Foo {
	*f += Foo(x)
	return f
}

func (f *Foo) Show() *Foo {
	fmt.Println(*f)
	return f
}

func main() {
	var i Foo = 12
	i.Add(10).Show().
		Add(10).Show().
		Add(10).Show()

	var baz mod.Baz = mod.Baz{}
	baz.Hi()
	baz.Num = 12
	fmt.Printf("%+v\n", baz)
}
