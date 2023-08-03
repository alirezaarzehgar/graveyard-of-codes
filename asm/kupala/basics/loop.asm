section .bss
	count	resb	2

section .text
	global _start
_start:
	mov	rax, 0
	mov	rdi, 0
	mov	rsi, count
	mov	rdx, 2
	syscall

	mov	r10b, [count]
	sub	r10b, 48
_loop:
	mov	[count], r10b
	mov	al, 48
	add	[count], al

	mov	rax, 1
	mov	rdi, 1
	mov	rsi, count
	mov	rdx, 2
	syscall

	dec	r10b
	jnz	_loop

	mov	rax, 60
	mov	rdi, 0
	syscall
