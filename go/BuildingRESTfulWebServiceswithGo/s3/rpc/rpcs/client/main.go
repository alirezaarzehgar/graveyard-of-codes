package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	c, e := rpc.Dial("tcp", "localhost:1234")
	if e != nil {
		log.Fatal("rpc.dial:", e)
	}
	defer c.Close()

	var reply int64
	e = c.Call("TimeServer.GetTime", struct{}{}, &reply)
	if e != nil {
		log.Println("rpc call:", e)
	}
	fmt.Println("Epoch time: ", reply)
}
