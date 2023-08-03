package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

func main() {
	os.Remove("data.gob")
	f, err := os.Create("data.gob")
	if err != nil {
		log.Fatal(err)
	}
	encoder := gob.NewEncoder(f)
	encoder.Encode(map[string]any{"name": "ali", "key": "value", "age": 18})
	f.Close()

	f, err = os.Open("data.gob")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var data map[string]any
	decoder := gob.NewDecoder(f)
	if err := decoder.Decode(&data); err != nil {
		log.Fatal(err)
	}
	fmt.Println("data:", data)
}
