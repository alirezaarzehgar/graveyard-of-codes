package main

import (
	"context"
	protoapi "grpc1/ptotoapi"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type MathServer struct {
	protoapi.UnimplementedMathServer
}

func (MathServer) Add(ctx context.Context, n *protoapi.NumberArgs) (*protoapi.Number, error) {
	return &protoapi.Number{Value: n.A + n.B}, nil
}

func (MathServer) Minus(ctx context.Context, n *protoapi.NumberArgs) (*protoapi.Number, error) {
	return &protoapi.Number{Value: n.A - n.B}, nil
}

func (MathServer) Div(ctx context.Context, n *protoapi.NumberArgs) (*protoapi.Number, error) {
	return &protoapi.Number{Value: n.A / n.B}, nil
}

func (MathServer) Mul(ctx context.Context, n *protoapi.NumberArgs) (*protoapi.Number, error) {
	return &protoapi.Number{Value: n.A * n.B}, nil
}

func main() {
	server := grpc.NewServer()
	protoapi.RegisterMathServer(server, MathServer{})
	reflection.Register(server)
	lister, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	server.Serve(lister)
}
