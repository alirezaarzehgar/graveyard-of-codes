package main

import (
	"fmt"
	"sort"
)

const (
	C1 = (iota + 1) * 2
	C2
	C3
	C4
	C5
)

func main() {
	il := []int{5, 3, 4, 3, 1, 3, 5, 7, 5, 6, 3, 4}

	fmt.Println(il)
	sort.Slice(il, func(i, j int) bool {
		return il[i] > il[j]
	})
	fmt.Println(il)
	sort.Slice(il, func(i, j int) bool {
		return il[i] < il[j]
	})
	fmt.Println(il)

	ml := map[string]any{
		"hi": 12,
	}
	_, ok := ml["hi"]
	fmt.Println(ok)
	_, ok = ml["hoy"]
	fmt.Println(ok)

	fmt.Println(C3)
}
