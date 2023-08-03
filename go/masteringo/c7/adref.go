package main

import (
	"fmt"
	"os"
	"reflect"
)

func printMethods(v any) {
	r := reflect.ValueOf(v).Elem()
	t := r.Type()

	for i := 0; i < r.NumMethod(); i++ {
		fmt.Println(t.Method(i).Name, "-->", r.Method(i).Type())
	}
}

type Task struct{}

func (t Task) Job1() {
	fmt.Println("Hello job 1")
}

func (t Task) Job2() {
	fmt.Println("Hello job 2")
}

func (t Task) Job3() {
	fmt.Println("Hello job 3")
}

func (t Task) Job4() {
	fmt.Println("Hello job 4")
}

func (t Task) Job5() {
	fmt.Println("Hello job 5")
}

func start(jobs []string) {
	for _, job := range jobs {
		method := reflect.ValueOf(&Task{}).MethodByName(job)
		method.Call(nil)
	}
}

func main() {
	printMethods(&os.Stdout)

	printMethods(&Task{})
	start([]string{"Job3", "Job1", "Job5"})
}
