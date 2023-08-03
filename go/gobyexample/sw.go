package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type Shape interface {
	InputData() error
	Area() float64
	Perimeter() float64
}

func getFloat(text string) (float64, error) {
	scnr := bufio.NewScanner(os.Stdin)

	fmt.Print(text)
	if !scnr.Scan() {
		return 0, scnr.Err()
	}

	n, err := strconv.ParseFloat(scnr.Text(), 64)
	if err != nil {
		return 0, errors.New("Failed to parse:" + err.Error())
	}

	if n <= 0 {
		return 0, errors.New(fmt.Sprint("Invalid x: ", n))
	}
	return n, nil
}

type Square struct {
	x float64
}

func (s *Square) InputData() error {
	x, err := getFloat("Enter X> ")
	if err != nil {
		return err
	}

	s.x = x
	return nil
}

func (s Square) Area() float64 {
	return s.x * s.x
}

func (s Square) Perimeter() float64 {
	return s.x * 4
}

type Circle struct {
	r float64
}

func (c *Circle) InputData() error {
	r, err := getFloat("Enter R> ")
	if err != nil {
		return err
	}

	c.r = r
	return nil
}

func (c Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.r
}

type Rectangle struct {
	x, y float64
}

func (r *Rectangle) InputData() error {
	x, err := getFloat("Enter X> ")
	if err != nil {
		return err
	}

	y, err := getFloat("Enter y> ")
	if err != nil {
		return err
	}

	r.x = x
	r.y = y
	return nil
}

func (r Rectangle) Area() float64 {
	return r.x * r.y
}

func (r Rectangle) Perimeter() float64 {
	return (r.x + r.y) * 2
}

func main() {
	fmt.Println(`
[1]- Square
[2]- Circle
[3]- Rectangle
`)
	ch, err := getFloat("Chose [1-3]> ")
	if err != nil {
		log.Fatal(err)
	}

	if ch < 1 || ch > 3 {
		log.Fatal("Invalid chose")
	}

	var shape Shape = []Shape{&Square{}, &Circle{}, &Rectangle{}}[int(ch-1)]
	shape.InputData()
	fmt.Println("Area:", shape.Area())
	fmt.Println("Perimeter:", shape.Perimeter())

	if v, ok := shape.(*Square); ok {
		fmt.Printf("Yeah is %#v\n", v)
	}

	if v, ok := shape.(*Circle); ok {
		fmt.Printf("Yeah is %#v\n", v)
	}

	if v, ok := shape.(*Rectangle); ok {
		fmt.Printf("Yeah is %#v\n", v)
	}

	func(s any) {
		switch v := s.(type) {
		case Square:
			fmt.Printf("fucking square: %#v\n", v)
		case Circle:
			fmt.Printf("fucking circle: %#v\n", v)
		case Rectangle:
			fmt.Printf("fucking rectangle: %#v\n", v)
		default:
			fmt.Printf("Unknown type: %#v\n", v)
		}
	}(shape)
}
