package main

import (
	"context"
	"fmt"
	protoapi "grpc1/ptotoapi"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := protoapi.NewMathClient(conn)

	fmt.Println(client.Add(context.TODO(), &protoapi.NumberArgs{A: 3, B: 7}))
	fmt.Println(client.Mul(context.TODO(), &protoapi.NumberArgs{A: 3, B: 7}))
	fmt.Println(client.Div(context.TODO(), &protoapi.NumberArgs{A: 30, B: 7}))
	fmt.Println(client.Minus(context.TODO(), &protoapi.NumberArgs{A: 23, B: 7}))
}
