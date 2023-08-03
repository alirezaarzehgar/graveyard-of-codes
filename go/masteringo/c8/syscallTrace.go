package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
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
	before := true
	for {
		if before {

			err = syscall.PtraceGetRegs(wpid, &r)
			if err != nil {
				break
			}

			fmt.Printf("syscall: %d\n", r.Orig_rax)
		}

		err := syscall.PtraceSyscall(wpid, 0)
		if err != nil {
			fmt.Println("PtraceSyscall:", err)
			return
		}
		_, err = syscall.Wait4(wpid, nil, 0, nil)
		if err != nil {
			fmt.Println("Wait4:", err)
			return
		}
		before = !before
	}
}
