package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide a filename!")
		return
	}

	fileData, err := os.ReadFile(args[1])
	if err != nil {
		log.Fatal(err)
	}

	var mdata map[string]any
	json.Unmarshal(fileData, &mdata)
	for key, value := range mdata {
		fmt.Println(key, value)
	}

	str := "Hello, World!"
	fmt.Printf("%x\n", str)

	h := make([]byte, len(str)*2)
	hex.Encode(h, []byte(str))
	fmt.Println(h)

	fmt.Println(hex.EncodeToString([]byte(str)))

	mdata["newkey"] = "hello world"
	fileData, err = json.Marshal(mdata)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(fileData))
}
