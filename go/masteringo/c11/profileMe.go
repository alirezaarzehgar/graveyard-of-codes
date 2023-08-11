package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func F1() {
	x := 2
	for i := 0; i < 52000; i++ {
		x = x * 2
		time.Sleep(622)
	}
}

func F2() {
	k := 0
	for i := 0; i < 4002; i++ {
		for j := 0; j < 12345; j++ {
			k++
			time.Sleep(1)
		}
	}
}

func F3() {
	var list []any
	for i := 0; i < 10002; i++ {
		list = append(list, fmt.Sprint("hi", i))
		time.Sleep(3223)
	}
}

func F4() {
	for i := 0; i < 5023; i++ {
		for j := 0; j < 80; j++ {
			for k := 0; k < 70; k++ {
				OK := 0
				OK++
				OK--
			}
		}
		time.Sleep(1000)
	}
}

func F5() {
	for i := 0; i < 1072; i++ {
		for i := 0; i < 34785; i++ {
			_ = fmt.Sprint("asfasf")
		}

		for i := 0; i < 17215; i++ {
			ok := ""
			_ = ok
		}
	}
}

func F6() {
	for i := 0; i < 5; i++ {
		F1()
	}
}

func main() {
	cpuFile, err := os.Create("/tmp/cpuProfile.out")
	if err != nil {
		log.Fatal(err)
	}
	defer cpuFile.Close()
	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	for i := 0; i < 5; i++ {
		F1()
	}
	fmt.Println("F1")
	for i := 0; i < 5; i++ {
		F2()
	}
	fmt.Println("F2")
	for i := 0; i < 5; i++ {
		F3()
	}
	fmt.Println("F3")
	for i := 0; i < 5; i++ {
		F4()
	}
	fmt.Println("F4")
	for i := 0; i < 5; i++ {
		F5()
	}
	fmt.Println("F5")
	for i := 0; i < 5; i++ {
		F6()
	}
	fmt.Println("F6")

	memory, err := os.Create("/tmp/memoryProfile.out")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("memory")
	for i := 0; i < 500; i++ {
		s := make([]byte, 500000000)
		if s == nil {
			log.Println("operation failed")
		}
	}
	if err = pprof.WriteHeapProfile(memory); err != nil {
		log.Fatal(err)
	}

}
