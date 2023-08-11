package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"
	"strings"
)

func main() {
	c, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer c.Close()

	for {
		var reply int
		fmt.Print(">> ")
		line, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			break
		}
		tokens := strings.Fields(line)
		if len(tokens) != 3 {
			fmt.Println("[CMD] [X] [Y]")
			continue
		}
		x, _ := strconv.Atoi(tokens[1])
		y, _ := strconv.Atoi(tokens[2])

		if err = c.Call("MyMath."+tokens[0], MyNums{x, y}, &reply); err != nil {
			log.Println(err)
			continue
		}
		fmt.Println("reply:", reply)
	}
}
