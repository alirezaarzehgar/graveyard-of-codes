package main

import (
	"log"
	"net"
	"net/http"
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
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("network listen: ", e)
	}

	http.Serve(l, nil)
}
