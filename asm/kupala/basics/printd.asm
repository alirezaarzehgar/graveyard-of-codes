section .bss
	dstr	resb	10

section .text
	global _start
_start:
	mov	al, 48
	add	al, 2
	mov	[dstr], al

	mov	al, 10
	mov	[dstr+1], al

	mov	rax, 1
	mov	rdi, 1
	mov	rsi, dstr
	mov	rdx, 2
	syscall

	mov	rax, 60
	mov	rdi, 0
	syscall

