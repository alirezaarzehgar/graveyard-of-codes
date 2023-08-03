package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strconv"
)

const TMP = "/tmp/rumination.tmp"

func main() {
	if len(os.Args) == 2 && os.Args[1] == "clear" {
		os.WriteFile(TMP, []byte("0"), fs.ModePerm)
		fmt.Println("Repeat:", 0)
		return
	}

	var num uint64 = 0
	if _, err := os.Stat(TMP); errors.Is(err, os.ErrNotExist) {
		os.WriteFile(TMP, []byte("0"), fs.ModePerm)
	} else {
		var err error

		data, err := os.ReadFile(TMP)
		if err != nil {
			log.Fatal(err)
		}

		num, err = strconv.ParseUint(string(data), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
	}

	num++
	fmt.Println("Repeat:", num)
	os.WriteFile(TMP, []byte(fmt.Sprint(num)), fs.ModePerm)
}
