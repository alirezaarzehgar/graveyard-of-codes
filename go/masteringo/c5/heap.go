package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	h := &MinHeap{1, 4, 3, 7, 5, 4, 8, 6, 5}
	heap.Init(h)
	fmt.Println(h.Len())
	fmt.Println(h)

	for h.Len() >= 1 {
		fmt.Println(h.Pop(), h)
	}

	for i := 0; i < 120; i += 3 {
		h.Push(i)
	}
	fmt.Println(h)
}
