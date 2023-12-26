package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func GetDirSize(path string) (int64, error) {
	contents, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var total int64
	for _, entry := range contents {
		if entry.IsDir() {
			temp, err := GetDirSize(filepath.Join(path, entry.Name()))
			if err != nil {
				log.Println(err)
				return 0, nil
			}
			total += temp
		} else {
			info, err := entry.Info()
			if err != nil {
				log.Println(err)
				return 0, err
			}
			total += info.Size()
		}
	}
	return total, nil
}

func main() {
	fmt.Println(GetDirSize(".."))
}
