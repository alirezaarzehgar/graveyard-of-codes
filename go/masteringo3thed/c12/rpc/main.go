package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"
)

type Opt struct{}

func (Opt) Func(args int, ret *int) error {
	fmt.Println("Hello, World")
	return nil
}

func main() {
	go func() {
		rpc.Register(new(Opt))
		l, err := net.Listen("tcp", ":1234")
		if err != nil {
			log.Fatal(err)
		}

		for {
			c, err := l.Accept()
			if err != nil {
				log.Println(err)
			}
			rpc.ServeConn(c)
		}
	}()

	time.Sleep(time.Second)
	c, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	c.Call("Opt.Func", 0, nil)
}
