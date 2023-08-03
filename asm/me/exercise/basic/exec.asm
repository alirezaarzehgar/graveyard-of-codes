%include "linux64.inc"

section .bss
	num	resb	1
section .text
	global _start
_start:
	pop	rcx

	pop	rdi
	pop	rdi
	sub	rsp, 8

	mov	rax, sys_execve
	mov	rsi, rsp
	mov	rdx, 0
	syscall

	exit 0
