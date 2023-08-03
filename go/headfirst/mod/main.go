package main

import (
	"syscall"

	"example.com/mod/lib"
	"example.com/mod/web"
)

func main() {
	pid, _, _ := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	if pid == 0 {
		web.HelloWeb(":8080")
	} else {
		web.Bobo(":9090")
	}
	lib.HelloWorld()
}
