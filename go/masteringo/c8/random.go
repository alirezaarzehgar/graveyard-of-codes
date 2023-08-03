package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

func getRandom() (any, error) {
	rf, err := os.Open("/dev/random")
	if err != nil {
		return nil, err
	}
	defer rf.Close()

	var v [10]byte
	binary.Read(rf, binary.LittleEndian, &v)

	return v, nil
}

func main() {
	v, err := getRandom()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("v:%+v\n", v)
}
