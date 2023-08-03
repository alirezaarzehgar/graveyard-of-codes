package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"syscall"
)

func main() {
	line, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		log.Fatal(err)
	}
	fmt.Print("line: ", line)

	fmt.Println("mamad" + string(97) + "ali")
	err = syscall.Access("/tmp/file", syscall.F_OK)
	if err == nil {
		fmt.Println("file exists")
	}

	f, err := os.OpenFile("/tmp/file", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}

	f.Write([]byte("Hello, World\n"))

	r, _ := http.Get("https://en.wikipedia.org/w/api.php?action=query&prop=revisions&titles=Pet_door&rvprop=content&format=json")

	fmt.Print(r.Body)
}
