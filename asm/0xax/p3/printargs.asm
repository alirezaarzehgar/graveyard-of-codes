section .text
	global _start

_start:
	pop	rcx
	cmp	rcx, 2
	jne	exit

	mov	rax, 1
	mov	rdi, 1
	add	rsp, 8
	pop	rsi
	mov	rdx, 2
	syscall

exit:
	mov	rax, 60
	mov	rdi, 0
	syscall
