package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs)
	go func() {
		for {
			sig := <-sigs
			fmt.Println("Bad signal:", sig)
		}
	}()

	select {}
}
