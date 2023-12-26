package main

import (
	"log"
	"net"
	"net/rpc"
	"time"
)

type TimeServer struct{}

func (t *TimeServer) GetTime(args struct{}, reply *int64) error {
	*reply = time.Now().Unix()
	return nil
}

func main() {
	rpc.Register(new(TimeServer))

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("network error:", e)
	}
	defer l.Close()

	for {
		c, e := l.Accept()
		if e != nil {
			log.Println("accept:", e)
			continue
		}

		rpc.ServeConn(c)
		c.Close()
	}
}
