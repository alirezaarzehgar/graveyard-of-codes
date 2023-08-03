package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	euid, _, _ := syscall.Syscall(syscall.SYS_GETEUID, 0, 0, 0)
	fmt.Println("My euid", euid)

	pid, _, _ := syscall.Syscall(syscall.SYS_GETPID, 0, 0, 0)
	fmt.Println("My pid", pid)

	syscall.Write(syscall.Stdout, []byte("Hello, World\n"))

	syscall.Exec("/bin/ls", []string{"ls", "-ltrh"}, os.Environ())
}
