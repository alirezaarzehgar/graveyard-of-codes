package person

import (
	"fmt"
)

type Person struct {
	age  int
	name string
}

func init() {
	fmt.Println("init")
}

func NewPerson(age int, name string) Person {
	return Person{age: age, name: name}
}

func (p *Person) Show() {
	fmt.Println("Age:", p.age)
	fmt.Println("Name:", p.name)
}

type Boy struct {
	Person
	dickSize int
}

func NewBoy(age int, name string, dickSize int) Boy {
	return Boy{
		Person:   Person{age: age, name: name},
		dickSize: dickSize,
	}
}

func (b *Boy) SetDick(x int) {
	b.dickSize = x
}

func (b Boy) DickSize() int {
	return b.dickSize
}
