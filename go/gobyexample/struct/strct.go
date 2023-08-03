package main

import (
	"github.com/alirezaarzehgar/hi/person"
)

func main() {
	var p person.Person = person.NewPerson(12, "ali")
	p.Show()

	b := person.NewBoy(44, "mamad", 18)
	b.SetDick(23)
	b.Show()
	println("Dick size:", b.DickSize())
}
