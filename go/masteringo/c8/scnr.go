package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scnr := bufio.NewScanner(os.Stdin)
	scnr.Scan()
	fmt.Println(scnr.Text())
}
