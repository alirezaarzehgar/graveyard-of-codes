package main

import "fmt"

const SIZE = 10

type Node struct {
	Value int
	Next  *Node
}

type HashTable struct {
	Table map[int]*Node
	Size  int
}

func hashFunction(i, size int) int {
	return i % size
}

func insert(hash *HashTable, value int) int {
	index := hashFunction(value, hash.Size)
	element := Node{Value: value, Next: hash.Table[index]}
	hash.Table[index] = &element
	return index
}

func hashTraverse(hash *HashTable) {
	for k := range hash.Table {
		if hash.Table[k] == nil {
			continue
		}

		t := hash.Table[k]
		fmt.Printf("{%d}", k)
		for t != nil {
			fmt.Print(" -> ", t.Value)
			t = t.Next
		}
		fmt.Println()
	}
}

func main() {
	table := make(map[int]*Node, SIZE)
	hash := &HashTable{Table: table, Size: SIZE}

	for i := 0; i < 120; i++ {
		insert(hash, i)
	}
	hashTraverse(hash)
}
