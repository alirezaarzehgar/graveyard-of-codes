section .data
	msg	db	"Hello, World!", 10
	msglen	equ	14
	loopcnt	equ	12

section .text
	global _start

_start:
	mov	r10, loopcnt

loop:
	mov	rax, 1
	mov	rdi, 1
	mov	rsi, msg
	mov	rdx, msglen
	syscall

	dec	r10
	jnz	loop

	mov	rax, 60
	mov	rdi, 0
	syscall
