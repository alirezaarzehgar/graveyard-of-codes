package main

import (
	"fmt"
	"sort"
	"strings"
)

type StrData []string

func (a StrData) Len() int           { return len(a) }
func (a StrData) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a StrData) Less(i, j int) bool { return strings.Compare(a[i], a[j]) < 0 }

func main() {
	d := StrData{"hi", "ali", "mohammad", "1", "32", "100"}
	sort.Sort(d)
	fmt.Println(d)
	sort.Sort(sort.Reverse(d))
	fmt.Println(d)
}
