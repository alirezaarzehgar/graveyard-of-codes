package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("pid:", os.Getpid())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGALRM)
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				fmt.Println("Caught:", sig)
			case syscall.SIGALRM:
				fmt.Println("WTF sig:", sig)
				return
			}
		}
	}()

	for {
		fmt.Print(".")
		time.Sleep(20 * time.Second)
	}
}
