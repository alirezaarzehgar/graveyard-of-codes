package main

import (
	"context"
	"grpc/datafiles"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := datafiles.NewMoneyTransactionClient(conn)

	r, err := c.MakeTransaction(context.Background(), &datafiles.TransactionRequest{From: "123", To: "312", Amount: 199.99})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(r)
}
