package main

type MyNums struct {
	X, Y int
}

type MyInterface interface {
	Add(args *MyNums, reply *int) error
	Minus(args *MyNums, reply *int) error
	Multiply(args *MyNums, reply *int) error
	Division(args *MyNums, reply *int) error
	Power(args *MyNums, reply *int) error
}
