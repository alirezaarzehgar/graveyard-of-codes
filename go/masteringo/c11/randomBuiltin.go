package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing/quick"
	"time"
)

func main() {
	type Point3D struct {
		X, Y, Z int
		S       float64
	}
	ran := rand.New(rand.NewSource(time.Now().UnixNano()))
	myValues := reflect.TypeOf(Point3D{})
	x, _ := quick.Value(myValues, ran)
	fmt.Println(x)
}
