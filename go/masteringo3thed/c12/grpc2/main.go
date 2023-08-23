package main

import (
	"context"
	"fmt"
	"grpc2/api"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type OptServer struct {
	api.UnimplementedOptServer
}

func (OptServer) Func(ctx context.Context, e *api.Empty) (*api.Empty, error) {
	fmt.Println("Hello, World")
	return &api.Empty{}, nil
}

func main() {
	go func() {
		server := grpc.NewServer()
		api.RegisterOptServer(server, OptServer{})
		reflection.Register(server)
		l, err := net.Listen("tcp", ":1234")
		if err != nil {
			log.Fatal(err)
		}
		server.Serve(l)
	}()

	conn, err := grpc.Dial("localhost:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	c := api.NewOptClient(conn)
	c.Func(context.TODO(), &api.Empty{})
}
