package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	target := rand.Intn(100) + 1

	for true {
		fmt.Print("Guess number: ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		guess, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Opps. Your guess was LOW")
		} else if guess > target {
			fmt.Println("Opps. Your guess was HIGH")
		} else {
			fmt.Println("Yo yo! You win!!")
			break
		}
	}
}
