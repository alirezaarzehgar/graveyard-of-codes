package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

func main() {
	if len(os.Args) != 2 {
		log.Printf("%s [address]\n", os.Args[0])
		os.Exit(1)
	}

	con, err := net.Dial("tcp", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	// sender
	go func() {
		stdin := bufio.NewReader(os.Stdin)
		for {
			fmt.Print(">> ")
			line, err := stdin.ReadString('\n')
			if err != nil {
				log.Println(err)
				continue
			}
			fmt.Fprint(con, line)
		}
	}()

	// reciever
	go func() {
		defer wg.Done()
		server := bufio.NewReader(con)
		for {
			msg, err := server.ReadString('\n')
			if err != nil {
				log.Println(err)
				continue
			}
			fmt.Print("server: ", msg)
			if strings.ToUpper(strings.TrimSpace(msg)) == "STOP" {
				con.Close()
				return
			}
		}
	}()

	wg.Wait()
	fmt.Println("Done everything!")
}
