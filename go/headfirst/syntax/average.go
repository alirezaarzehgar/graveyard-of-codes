package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal(os.Args[0], " [filename]")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	var sum float64 = 0
	scnr := bufio.NewScanner(file)
	for scnr.Scan() {
		line := strings.TrimSpace(scnr.Text())
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			log.Println("Invalid floating-point value: ", num)
			continue
		}

		sum += num
		count++
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scnr.Err() != nil {
		log.Fatal(scnr.Err())
	}

	fmt.Println("sum: ", sum)
	fmt.Println("average: ", sum/float64(count))

	a := []int{1, 2, 3}
	b := [3]int{1, 2, 3}

	a = append(a, 2)

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))

	fmt.Println(a[1:3], b[1:3])

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var as, ds []int

	as = append(as, arr[1:4]...)
	ds = append(ds, arr[4:8]...)

	arr[4] = 22

	fmt.Println(arr)
	fmt.Println(as)
	fmt.Println(ds)
}
