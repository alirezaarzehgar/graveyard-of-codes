section .data
	l17str	db	"Enter your name: "
	l6str	db	"Name: "

section .bss
	name	resb	10

section .text
	global _start

_start:
	mov	rax, 1
	mov	rdi, 1
	mov	rsi, l17str
	mov	rdx, 17
	syscall

	mov	rax, 1
	mov	rdi, 1
	mov	rsi, l6str
	mov	rdx, 6
	syscall

	mov	rax, 0
	mov	rdi, 0
	mov	rsi, name
	mov	rdx, 10
	syscall

	mov	rax, 1
	mov	rdi, 1
	mov	rsi, name
	mov	rdx, 10
	syscall

	mov	rax, 60
	mov	rdi, 0
	syscall
