package main

import (
	"context"
	"fmt"
	"time"
)

func doAnything(ctx context.Context) {
	go func() {
		time.Sleep(5)
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Fucking done:", ctx.Err())
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	doAnything(ctx)
	// context.TODO()
	// context.Background()
}
