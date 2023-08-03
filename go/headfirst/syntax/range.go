package main

func main() {
	var nums []int = []int{1, 5, 2, 7, 43, 5, 3, 4}
	var sum int

	for _, v := range nums {
		sum += v
	}

	println("sum:", sum)
	println("average:", sum/len(nums))
}
