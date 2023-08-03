package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Pass me a command!")
		return
	}

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{Ptrace: true}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	err := cmd.Wait()
	fmt.Printf("state: %v\n", err)
	wpid := cmd.Process.Pid

	var r syscall.PtraceRegs
	if err = syscall.PtraceGetRegs(wpid, &r); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Registers: %#v\n", r)
	fmt.Printf("R15=%d, Gs=%d\n", r.R15, r.Gs)

	time.Sleep(time.Second * 2)
}
