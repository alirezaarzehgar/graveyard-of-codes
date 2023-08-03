package main

import (
	"fmt"
	"unsafe"
)

type data struct {
	a int
	b int
	c int
}

func main() {
	var list = [...]int{1, 2, 3, 4, 5}
	memaddr := uintptr(unsafe.Pointer(&list))

	for i := 0; i < len(list); i++ {
		ptr := (*int)(unsafe.Pointer(memaddr))
		fmt.Print(*ptr, " ")
		memaddr = uintptr(unsafe.Pointer(ptr)) + unsafe.Sizeof(list[0])
	}
	fmt.Println()

	var num uintptr = 10000000000000
	var c *int16 = (*int16)(unsafe.Pointer(&num))
	fmt.Printf("%d\n", *c)
}
