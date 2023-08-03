package main

import (
	"container/list"
	"fmt"
)

func main() {
	values := list.New()

	e1 := values.PushBack("One")
	e2 := values.PushBack("Two")
	values.PushFront("Three")
	values.InsertBefore("Four", e1)
	values.InsertAfter("Five", e2)
	values.Remove(e1)
	values.Remove(e2)

	values.InsertAfter("FiveFive", e2)
	values.PushBackList(values)

	for t := values.Back(); t != nil; t = t.Prev() {
		fmt.Print(t.Value, " ")
	}
	fmt.Println()

	for t := values.Front(); t != nil; t = t.Next() {
		fmt.Print(t.Value, " ")
	}
	fmt.Println()

	values.Init()
}
