package main

import (
	"fmt"
)

func showG[T any](x T) {
	fmt.Println(x)
}

func sumNum[V int | float64](numbers []V) V {
	var sum V
	for _, v := range numbers {
		sum += v
	}
	return sum
}

type Number interface {
	int | float64
}

func sumNumber[V Number](numbers []V) V {
	var sum V
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func main() {
	showG[int](1)

	fmt.Println(
		sumNum[int]([]int{1, 2, 3, 44}),
		sumNum([]float64{1, 2, 3, 44}),
	)
}
