package main

import (
	"context"
	"fmt"
	"grpc/app"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	app.UnimplementedMessageServiceServer
}

func (server) SayIt(ctx context.Context, r *app.Request) (*app.Response, error) {
	fmt.Println("Text:", r.Text)
	fmt.Println("Subtext:", r.Subtext)

	response := &app.Response{
		Text:    r.Text,
		Subtext: "Got it!",
	}

	return response, nil
}

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	app.RegisterMessageServiceServer(s, server{})
	s.Serve(l)
}
