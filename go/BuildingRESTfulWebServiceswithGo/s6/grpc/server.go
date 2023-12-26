package main

import (
	"context"
	"grpc/datafiles"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	datafiles.UnimplementedMoneyTransactionServer
}

func (server) MakeTransaction(ctx context.Context, r *datafiles.TransactionRequest) (*datafiles.TransactionResponse, error) {
	log.Println("do somethings")
	return &datafiles.TransactionResponse{Confirmation: true}, nil
}

func main() {
	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	datafiles.RegisterMoneyTransactionServer(s, &server{})
	reflection.Register(s)
	s.Serve(l)
}
