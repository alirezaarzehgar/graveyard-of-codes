package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var data any
	raw, _ := os.ReadFile("data.json")

	json.Unmarshal(raw, &data)

	for k, v := range data.(map[string]any) {
		fmt.Println(k, ":")

		_ = v
		for _, v := range v.([]any) {
			fmt.Println("\t", v)
		}
	}
}
