package main

import (
	"context"
	"fmt"
	"grpc/app"
	"log"

	"google.golang.org/grpc"
)

func AboutToSayIt(ctx context.Context, a app.MessageServiceClient, text string) (*app.Response, error) {
	request := &app.Request{
		Text:    text,
		Subtext: "New Message!",
	}
	r, err := a.SayIt(ctx, request)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := app.NewMessageServiceClient(conn)
	// r, err := AboutToSayIt(context.Background(), client, "My Message!")
	r, err := client.SayIt(context.Background(), &app.Request{Text: "My Message!", Subtext: "New Message!"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Text:", r.Text)
	fmt.Println("Subtext:", r.Subtext)
}
