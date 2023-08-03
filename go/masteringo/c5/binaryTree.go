package main

import "fmt"

type Tree struct {
	Right *Tree
	Value int
	Left  *Tree
}

func traverse(t *Tree, depth int) {
	if t == nil {
		return
	}

	traverse(t.Left, depth+1)
	for i := 0; i < depth; i++ {
		fmt.Print("\t")
	}
	fmt.Println(t.Value)
	traverse(t.Right, depth+1)
}

func main() {
	t := Tree{
		Right: &Tree{
			Right: &Tree{Value: 4},
			Value: 2,
			Left:  &Tree{Value: 8},
		},
		Value: 1,
		Left: &Tree{
			Right: &Tree{Value: 9},
			Value: 3,
			Left:  &Tree{Value: 12},
		},
	}

	traverse(&t, 0)
}
