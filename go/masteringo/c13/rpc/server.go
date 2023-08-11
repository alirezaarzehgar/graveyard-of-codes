package main

import (
	"fmt"
	"log"
	"math"
	"net"
	"net/rpc"
)

type MyMath struct{}

func (MyMath) Add(args *MyNums, reply *int) error {
	*reply = args.X + args.Y
	return nil
}

func (MyMath) Minus(args *MyNums, reply *int) error {
	*reply = args.X - args.Y
	return nil

}
func (MyMath) Multiply(args *MyNums, reply *int) error {
	*reply = args.X * args.Y
	return nil
}

func (MyMath) Division(args *MyNums, reply *int) error {
	*reply = args.X / args.Y
	return nil
}

func (MyMath) Power(args *MyNums, reply *int) error {
	*reply = int(math.Pow(float64(args.X), float64(args.Y)))
	return nil
}

func main() {
	rpc.Register(new(MyMath))
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		defer c.Close()
		go func(c net.Conn) {
			fmt.Println(c.RemoteAddr())
			rpc.ServeConn(c)
		}(c)
	}
}
