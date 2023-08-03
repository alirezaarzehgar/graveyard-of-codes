package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Square struct {
	X float64
}

func (s Square) Area() float64 {
	return s.X * s.X
}

func (s Square) Perimeter() float64 {
	return s.X * 4
}

type Circle struct {
	R float64
}

func (c Circle) Area() float64 {
	return c.R * c.R * math.Pi
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}

func calculate(x Shape) {
	_, ok := x.(Square)
	if ok {
		fmt.Println("\033[31mShape is square:\033[0m")
	}

	_, ok = x.(Circle)
	if ok {
		fmt.Println("\033[31mShape is circle:\033[0m")
	}
patience
	fmt.Println("Area:", x.Area())
	fmt.Println("Perimeter:", x.Perimeter())
}

func main() {
	calculate(Square{5})
	calculate(Circle{5})
}
