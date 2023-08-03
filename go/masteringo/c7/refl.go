package main

import (
	"errors"
	"fmt"
	"reflect"
)

type A struct {
	A int
	B string
}

type B struct {
	C float64
	D error
	E uint32
	F func() string
}

func showDetail(t any) {
	r := reflect.ValueOf(t).Elem()
	iType := r.Type()
	fmt.Println("type:", iType)
	fmt.Println("number:", r.NumField())

	for i := 0; i < r.NumField(); i++ {
		fmt.Print("Field name", iType.Field(i).Name)
		fmt.Print(" with type ", r.Field(i).Type())
		fmt.Println(" value", r.Field(i).Interface())
	}
	fmt.Println("-------------")
}

func main() {
	x := 10
	xRef := reflect.ValueOf(&x).Elem()
	xType := xRef.Type()
	fmt.Printf("the type of x is %s.\n", xType)

	fmt.Println("++++++++++")
	showDetail(&A{A: 14, B: "Hello, World"})
	showDetail(&B{
		C: 14.53, D: errors.New("err"), E: 444,
		F: func() string { return "hoyya" },
	})
}
